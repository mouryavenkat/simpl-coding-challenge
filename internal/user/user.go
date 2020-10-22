package user

import (
	"errors"
	"fmt"
	"log"
	"simpl-coding-challenge/internal"
	"simpl-coding-challenge/internal/infrastructure"
	"strings"
	"time"
)

func NewUser(dbClient infrastructure.MySQL) internal.User {
	return &userInfo{
		dbClient: dbClient,
	}
}

/*
	AddNewUser adds a new user to the simpl store.
	Constraints: UserName - Unique
*/
func (u *userInfo) AddNewUser(userName string, userEmail string, creditLimit int64) (message string, err error) {
	if createError := u.dbClient.Insert("insert into user_details values(?, ?, ?, ? )",
		userName, userEmail, creditLimit, 0); createError != nil {
		log.Println("Error: " + createError.Error())
		return message, createError
	}
	return fmt.Sprintf("userName(%d)", creditLimit), nil
}

func (u *userInfo) Purchase(userName string, merchantName string, transactionAmount int64) (message string, err error) {
	var userCreditBalance int64
	var merchantDiscount float64
	var dueAmount int64
	rows, err := u.dbClient.Query("select credit_limit - due_amount, due_amount from user_details where user_name = ?", userName)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		if err = rows.Scan(&userCreditBalance, &dueAmount); err != nil {
			return "", err
		}
	}

	rows, err = u.dbClient.Query("select merchant_discount from merchant_details where merchant_name = ?", merchantName)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		if err = rows.Scan(&merchantDiscount); err != nil {
			return "", err
		}
	}

	txn, err := u.dbClient.Begin()
	if err != nil {
		return "", err
	}
	if transactionAmount > userCreditBalance {
		return "", errors.New("credit Limit Exceeded")
	}
	if _, err = txn.Exec("insert into transactions(user_name, merchant_name, txn_value, merchant_discount) values(?,?,?,?)",
		userName, merchantName, transactionAmount, merchantDiscount); err != nil {
		txn.Rollback()
		return "", err
	}
	if _, err := txn.Exec("update user_details set due_amount = ? where user_name = ?", dueAmount+transactionAmount, userName); err != nil {
		txn.Rollback()
		return "", err
	}
	return "success!", txn.Commit()
}

func (u *userInfo) PayDueAmount(userName string, amount int64) error {
	txn, err := u.dbClient.Begin()
	if err != nil {
		return err
	}
	if _, err = txn.Exec("update user_details set due_amount=due_amount - ? where user_name = ?", amount, userName); err != nil {
		txn.Rollback()
		return err
	}
	// TODO: Generate unique payment ID
	if _, err := txn.Exec("insert into user_payments(user_name, amount_paid, paid_on, payment_id) values(?,?,?,?)", userName,
		amount, time.Now(), "random_id"); err != nil {
		txn.Rollback()
		return err
	}
	return txn.Commit()
}

func (u *userInfo) GetDueAmount(userNames ...string) (dueAmounts map[string]int64, err error) {
	var query string
	if len(userNames) > 0 {
		baseQuery := "select user_name, due_amount from user_details where user_name in ('%s') and due_amount > 0"
		query = fmt.Sprintf(baseQuery, strings.Join(userNames, "','"))
	} else {
		query = "select user_name, due_amount from user_details where due_amount > 0"
	}

	rows, err := u.dbClient.Query(query)
	if err != nil {
		return nil, err
	}
	dueAmounts = make(map[string]int64)
	for rows.Next() {
		var userName string
		var dueAmount int64
		if err = rows.Scan(&userName, &dueAmount); err != nil {
			return nil, err
		}
		dueAmounts[userName] = dueAmount
	}
	return dueAmounts, nil
}

func (u *userInfo) GetCreditLimitExceededUsers() (users []string, err error) {
	rows, err := u.dbClient.Query("select user_name from user_details where credit_limit - due_amount = ?", 0)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var userName string
		if err = rows.Scan(&userName); err != nil {
			return nil, err
		}
		users = append(users, userName)
	}
	return users, nil
}

func (u *userInfo) GetTotalUserDues() (dueAmount int64, err error) {
	rows, err := u.dbClient.Query("select sum(due_amount) from user_details")
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&dueAmount); err != nil {
			return 0, err
		}
		return dueAmount, nil
	}
	return 0, nil
}
