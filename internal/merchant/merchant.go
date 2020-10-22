package merchant

import (
	"fmt"
	"log"
	"simpl-coding-challenge/internal"
	"simpl-coding-challenge/internal/infrastructure"
)

func NewMerchant(dbClient infrastructure.MySQL) internal.Merchant {
	return &merchantInfo{
		dbClient: dbClient,
	}
}

func (m merchantInfo) AddNewMerchant(merchantName string, emailID string, merchantDiscount float64) (message string, err error) {
	if createError := m.dbClient.Insert("insert into merchant_details(merchant_name, email_id, merchant_discount) values(?,?,?)",
		merchantName, emailID, merchantDiscount); createError != nil {
		log.Println("Error: " + createError.Error())
		return message, createError
	}
	return fmt.Sprintf("%s(%.2f)", merchantName, merchantDiscount), nil
}

func (m merchantInfo) UpdateDiscountPercentage(merchantName string, merchantDiscount float64) (message string, err error) {
	if updateError := m.dbClient.Update("update merchant_details set merchant_discount = ? where merchant_name = ?", merchantDiscount, merchantName); updateError != nil {
		log.Println("Error: " + updateError.Error())
		return message, updateError
	}
	return fmt.Sprintf("%s(%.2f)", merchantName, merchantDiscount), nil
}

func (m merchantInfo) GetTotalMerchantDiscount(merchantName string) (merchantDiscount float64, err error) {
	rows, queryError := m.dbClient.Query("select sum(txn_value * merchant_discount/100) from transactions where merchant_name = ?", merchantName)
	if queryError != nil {
		log.Println("Error: " + queryError.Error())
		return 0, queryError
	}
	for rows.Next() {
		if err = rows.Scan(&merchantDiscount); err != nil {
			return 0, err
		}
		return merchantDiscount, nil
	}
	return 0, nil
}
