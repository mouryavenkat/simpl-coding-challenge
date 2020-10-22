package internal

type User interface {
	// AddNewUser adds a new user to the simpl store.
	AddNewUser(userName string, userEmail string, creditLimit int64) (message string, err error)
	// Purchase lets a user transact against a merchant
	Purchase(userName string, merchantName string, transactionAmount int64) (message string, err error)
	// PayDueAmount updates the amount user paid as part of this due amount and updates the credit limit
	PayDueAmount(userName string, amount int64) error
	// GetDueAmount return due amount of users if userName is not nil, else returns for all users
	GetDueAmount(userNames ...string) (userDuesMap map[string]int64, err error)
	// GetCreditLimitExceededUsers returns all the users info whose credit limit has exceeded.
	GetCreditLimitExceededUsers() (users []string, err error)
	// GetTotalUserDues returns the sum of total due amount for all the users
	GetTotalUserDues() (dueAmount int64, err error)
}
