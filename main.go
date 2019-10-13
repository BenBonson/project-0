package main

import (
	"fmt"
	"strings"

	//links to bank.go

	"database/sql"

	//imports the db

	"github.com/190930-UTA-CW-Go/project-0/arena"
	"github.com/190930-UTA-CW-Go/project-0/bank"
	"github.com/190930-UTA-CW-Go/project-0/name"
	"github.com/190930-UTA-CW-Go/project-0/shop"

	//"github.com/190930-UTA-CW-Go/project-0/navigation"
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

	//endless loop
	for {
		Nav()
	}

	//navigation.Nav()

	//check table
	//getAll(db)

	//Current location Home
	// var clocal string
	// clocal == "H"
	//add locations to go, shop arena bank

	//call the func bank
	//bank
	//location = bank or something would probably be cleaner
	//check for bank account

	//make a func that runs bank   TRY SWITCH STATMENTS

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

//Nav allows the player to choose their destination
func Nav() string {
	name := name.GetName()
	var tolocation string
	fmt.Println("Where would you like to go", name, "? \n [S]hop [A]rena [B]ank")
	fmt.Scanln(&tolocation)
	var tolocationlower = strings.ToLower(tolocation)

	switch tolocationlower {
	case "s":
		shop.Shop()
	case "a":
		arena.Arena()
	case "b":
		bank.Bank()
	default:
		fmt.Println("Sorry not avalible yet")
		// var tolocation string
		// fmt.Println("Sorry not an option please choose [S]hop [A]rena [B]ank")
		// fmt.Scanln(&tolocation)
		// var tolocationlower = strings.ToLower(tolocation)
	}
	return "nope" //Apparently I need a return but don't really want to use it
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
