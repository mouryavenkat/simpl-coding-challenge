package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"simpl-coding-challenge/internal/infrastructure/mysql"
	"simpl-coding-challenge/internal/merchant"
	"simpl-coding-challenge/internal/user"
	"strconv"
	"strings"
	"time"
)

func main() {
	log.Println("Started")
	db, err := sql.Open("mysql", "root:rootroot@tcp(mysqldb:3306)/simpl")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
		splitText := strings.Split(text, " ")
		if len(splitText) == 0 {
			fmt.Println("Enter proper command")
			continue
		}
		mysqlClient := mysql.NewSqlWrapper(db)
		userClient := user.NewUser(mysqlClient)
		merchantClient := merchant.NewMerchant(mysqlClient)
		switch splitText[0] {
		case "new":
			switch splitText[1] {
			case "user":
				userName := splitText[2]
				emailID := splitText[3]
				fmt.Println()
				creditLimit, err := strconv.ParseInt(splitText[4], 10, 64)
				if err != nil {
					fmt.Println("Error: Enter proper integer credit limit ", err.Error())
					continue
				}
				message, err := userClient.AddNewUser(userName, emailID, creditLimit)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				fmt.Println(message)
			case "merchant":
				userName := splitText[2]
				emailID := splitText[3]
				discountPercentage, err := strconv.ParseFloat(splitText[4], 10)
				if err != nil {
					fmt.Println("Error: Enter proper integer credit limit")
					continue
				}
				message, err := merchantClient.AddNewMerchant(userName, emailID, discountPercentage)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				fmt.Println(message)
			case "txn":
				userName := splitText[2]
				merchantName := splitText[3]
				transactionValue, err := strconv.ParseInt(splitText[4], 10, 64)
				if err != nil {
					fmt.Println("Error: Enter proper integer credit limit")
					continue
				}
				message, err := userClient.Purchase(userName, merchantName, transactionValue)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				fmt.Println(message)
			}
		case "report":
			switch splitText[1] {
			case "users-at-credit-limit":
				users, err := userClient.GetCreditLimitExceededUsers()
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				var output string
				for _, user := range users {
					output += user + "\n"
				}
				fmt.Println(output)
			case "discount":
				merchantName := splitText[2]
				discount, err := merchantClient.GetTotalMerchantDiscount(merchantName)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				fmt.Println("Total amount earned via merchant " + merchantName + "= " + strconv.FormatFloat(discount, 'f', -1, 64))
			case "total-dues":
				userNames := splitText[2:]
				userDues, err := userClient.GetDueAmount(userNames...)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				var output string
				for user, due := range userDues {
					output += user + " : " + strconv.FormatInt(due, 10) + "\n"
				}
				fmt.Println(output)
			}

		case "payback":
			userName := splitText[1]
			amount, err := strconv.ParseInt(splitText[2], 10, 64)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			err = userClient.PayDueAmount(userName, amount)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			fmt.Println("success!")
		case "set":
			switch splitText[1] {
			case "discount":
				merchantName := splitText[2]
				discountPercentage, err := strconv.ParseFloat(splitText[3], 10)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				message, err := merchantClient.UpdateDiscountPercentage(merchantName, discountPercentage)
				if err != nil {
					fmt.Println("Error: ", err.Error())
					continue
				}
				fmt.Println(message)
			}
		case "help":
			fmt.Println("1-new user <username> <email_id> <credit_limit>")
			fmt.Println("2-new merchant <merchant_name> <email_id> <discount_percentage>")
			fmt.Println("3-new txn <username> <merchant_name> <txn_value>")
			fmt.Println("4-report users-at-credit-limit")
			fmt.Println("5-report discount")
			fmt.Println("6-report total-dues [<space separated usernames>]")
			fmt.Println("7-payback <user_name> <amount>")
			fmt.Println("8-set discount <merchant_name> <newer_discount_percentage>")
		default:
			fmt.Println("Enter valid command ", text)
		}
	}

}
