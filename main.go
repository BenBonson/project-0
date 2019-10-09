package main

import (
	"fmt"

	//links to bank.go
	"github.com/190930-UTA-CW-Go/project-0/bank"

	"strings"

	"database/sql"

	//imports the db
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func main() {
	//gets data from db
	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", datasource)
	defer db.Close()
	if err != nil {
		panic(err)
	}
	getAll(db)
	//get player name
	var name string
	fmt.Println("Hello adventurer please state your name.")
	fmt.Scanln(&name)

	//check for bank account
	if bank.Account() {
		fmt.Println("Welcome back", name, "how can I help you?") //add options deposit withdraw
	} else {
		fmt.Println("Welcome", name, "I see you do not have an account with us would you like to set one up? [Y] or [N]")
	}

	//create account       find out how to update db with user input
	var acreate string
	fmt.Scanln(&acreate)
	var acreatelower = strings.ToLower(acreate)
	//check for user input
	//fmt.Println(acreatelower)
	if acreatelower == "y" {
		fmt.Println("Alright let me create that for you. Just so you know if you store your gold at the bank you will be unable to bet it however you wont lose it if you are beaten.")
		//figure out how to Serial aka add one to the id value
		db.Exec("INSERT INTO bankaccount VALUES (4, name, 10.3)")
		getAll(db)
	} else {
		fmt.Println("Well, good luck with that. Just so you know if you store your gold at the bank you will be unable to bet it however you wont lose it if you are beaten.")
	}

	//Notes and references

	// initbalance := flag.Float64("balance", 100, "default balance")
	// name := flag.String("name", "Mehrab", "default name")
	// flag.Parse()

	// fmt.Println(*name)
	// fund := fund.NewFund(*initbalance)

	// lock := make(chan bool)

	// go func() {
	// 	fund.Withdraw(50)
	// 	lock <- false
	// }()

	// //fmt.Println("Hit enter to continue")
	// //fmt.Scanln()

	// <-lock
	// fund.Withdraw(45)

	// fmt.Println(fund.Balance())

	//Current problems,
	//Serial the id
	//print out database
	//does the database stay changes or not
}

//attempt
func getAll(db *sql.DB) {
	rows, _ := db.Query("SELECT * FROM bankaccount")
	for rows.Next() {
		var id int
		var name string
		var funds float64
		rows.Scan(&id, &name, &funds)
		fmt.Println(id, name, funds)
	}
}
