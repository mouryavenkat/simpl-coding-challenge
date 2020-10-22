package internal

type Merchant interface {
	// AddNewMerchant onboards a new merchant to the simpl ecosystem
	AddNewMerchant(merchantName string, emailID string, merchantDiscount float64) (message string, err error)
	// UpdateDiscountPercentage allows merchant to update the discount percentage he/she provides to simpl per txn
	UpdateDiscountPercentage(merchantName string, merchantDiscount float64) (message string, err error)
	// GetTotalMerchantDiscount returns the discount that simpl received from a merchant till date
	GetTotalMerchantDiscount(merchantName string) (totalMerchantDiscount float64, err error)
}
