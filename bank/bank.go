package bank

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/190930-UTA-CW-Go/project-0/name"
	//"github.com/190930-UTA-CW-Go/project-0/navigation"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// Account bool Checks for account
func Account() bool {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	//for name in database scan through with range

	//scan through database for name
	// if {
	// 	name is in the database
	//	var account = true
	// } else {
	//	var account = false
	//}

	var account = false
	return account
}

//Bank starts when player enters the Bank
func Bank() {
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}

	name := name.GetName()

	if Account() {
		fmt.Println("Welcome back", name, "how can I help you? [W]ithdraw [D]eposit") //add options deposit withdraw
		var moneydo string
		fmt.Scanln(&moneydo)
		var moneydolower = strings.ToLower(moneydo)
		if moneydolower == "w" {
			fmt.Println("How much?")
		} else if moneydolower == "d" {
			fmt.Println("How much?")
		} else {
			fmt.Println("[W]ithdraw [D]eposit please")
		}
	} else {
		fmt.Println("Welcome", name, "I see you do not have an account with us would you like to set one up? [Y] or [N]")
		//create account
		var acreate string
		fmt.Scanln(&acreate)
		var acreatelower = strings.ToLower(acreate)
		//check for user input
		//fmt.Println(acreatelower)
		if acreatelower == "y" {
			//change bank account to true
			//bank.Account() account = true
			fmt.Println("Alright let me create that for you. Just so you know if you store your gold at the bank you will be unable to bet it however you wont lose it if you are beaten.")
			//add new account with a 0 balance under the player name
			db.Exec("insert into bankaccount (name, funds) values ($1, 10.5)", name)
			//check table
			//getAll(db)
			/*
				loop to case b again
			*/
		} else {
			fmt.Println("Well, good luck with that. Just so you know if you store your gold at the bank you will be unable to bet it \n however you wont lose it if you are beaten.")
			/*
				var tolocation string
				fmt.Println("Where would you like to go?. \n [S]hop [A]rena [B]ank")
				fmt.Scanln(&tolocation)
				var tolocationlower = strings.ToLower(tolocation)
			*/
		}
	}
}

// func d() bool {
//     var e bool
//     return e
// }

// if d() {
//     fmt.Printf("true")
// }

//notes and references

// // Fund is a simple struct representing a balance
// type Fund struct {
// 	balance float64
// }

// // NewFund is a basic constructor for Fund struct
// func NewFund(i float64) *Fund {
// 	return &Fund{
// 		balance: i,
// 	}
// }

// // Balance returns the current balance of a Fund
// func (f *Fund) Balance() float64 {
// 	return f.balance
// }

// // Withdraw removes an amount from the current balance of a Fund
// func (f *Fund) Withdraw(i float64) {
// 	f.balance -= i
// }
