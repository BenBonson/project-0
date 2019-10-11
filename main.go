package main

import (
	"fmt"

	//links to bank.go

	"strings"

	"database/sql"

	//imports the db
	"github.com/190930-UTA-CW-Go/project-0/bank"
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
	//check table
	//getAll(db)

	//Current location Home
	// var clocal string
	// clocal == "H"

	//get player name
	var name string
	fmt.Println("Hello adventurer please state your name.")
	fmt.Scanln(&name)
	//add locations to go, shop arena bank
	var tolocation string
	fmt.Println("Thanks for playing my game, where would you like to go?. \n [S]hop [A]rena [B]ank")
	fmt.Scanln(&tolocation)
	var tolocationlower = strings.ToLower(tolocation)
	//call the func bank
	//bank
	//location = bank or something would probably be cleaner
	//check for bank account

	//make a func that runs bank   TRY SWITCH STATMENTS
	switch tolocationlower {
	case "s":
		fmt.Println("Sorry not avalible yet")
	case "a":
		fmt.Println("Sorry not avalible yet")
	case "b":
		if bank.Account() {
			fmt.Println("Welcome back", name, "how can I help you?") //add options deposit withdraw
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
	default:
		//var tolocation string
		fmt.Println("Sorry not an option please choose [S]hop [A]rena [B]ank")
		// fmt.Scanln(&tolocation)
		// var tolocationlower = strings.ToLower(tolocation)

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
}

//getAll
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
