# Insert Project Name Here
## Insert Name Here
Insert project description here.

# User Stories
- [x] List
- [] Each
- [] User
- [] Story

# Instructions
Insert environment, build, and execution documentation here.


# Incase go deletes some of my imports
import (
	"fmt"

	//links to bank.go
	"github.com/190930-UTA-CW-Go/project-0/bank"

	"strings"

	"database/sql"

	//imports the db
	_ "github.com/lib/pq"
)

# Data Base Notes
cd db
docker build -t bankaccount .
docker run --name bankaccount -d -p 5432:5432 bankaccount
docker run --name bankaccount -it -p 5432:5432 bankaccount


after initial creation
docker exec -d bankaccount psql -U postgres
docker exec -it bankaccount psql -U postgres
docker start bankaccount
docker stop bankaccount
SELECT * FROM bankaccount;
\q

reset
cd into db
docker stop bankaccount
docker rm bankaccount
check with docker ps

# Some of the attempts and reference material used to recognise if the player already has an account

//Account bool Checks for account
// func Account(db *sql.DB, username string) bool {
// 	nameuser := name.GetName()
// 	sqlStatement := `SELECT name, funds FROM users WHERE name=$1;`
// 	var funds float64
// 	var name string
// 	var account bool
// 	// Replace 3 with an ID from your database or another random
// 	// value to test the no rows use case.
// 	row := db.QueryRow(sqlStatement, nameuser)
// 	//err := row.Scan(&name, &funds); err {
// 	// case sql.ErrNoRows:
// 	// 	fmt.Println("No rows were returned!")
// 	// if row := nil {
// 	if rowExists("SELECT id FROM feed_items WHERE url=$1", nameuser)
// 		var account true
// 		//fmt.Println(name, funds)
// 	} else {
// 		var account false
// 	}
// 	return account
// }

// func Account(query string, args ...interface{}) bool {
// 	var exists bool
// 	query = fmt.Sprintf("SELECT exists (%s)", query)
// 	err := db.QueryRow(query, args...).Scan(&exists)
// 	if err != nil && err != sql.ErrNoRows {
// 		glog.Fatalf("error checking if row exists '%s' %v", args, err)
// 	}
// 	return exists
// }


// if rows != nil {
	// 	return true
	// } else {
	// 	return false
	// }

// // Pokedex is a directory of Pokemon
// type Account struct {
// 	directory []bankaccount
// }

// // List will print out all Pokemon in the directory
// func (b *bankaccount) List() {
// 	for _, bankaccount := range b.directory {
// 		fmt.Println(bankaccount.name, bankaccount.funds)
// 	}
// }

//Account checks for player bank account
// func Account(name string) bool {
// 	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", datasource)
// 	defer db.Close()
// 	if err != nil {
// 		panic(err)
// 	}
// 	var row *sql.Row
// 	row = db.QueryRow("SELECT name FROM bankaccount WHERE name = $1", name)
//}

// Account bool Checks for account
// func Account() bool {
// 	datasource := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)
// 	db, err := sql.Open("postgres", datasource)
// 	defer db.Close()
// 	if err != nil {
// 		panic(err)
// 	}

// name := name.GetName()
//for name in database scan through with range
// for dbname, datasource := range datasource{
// 	if name
// }

// func UserExists(db * sql.DB, username string) bool {
// 	sqlStmt := `SELECT username FROM userinfo WHERE username = ?`
// 	err := db.QueryRow(sqlStmt, username).Scan(&username)
// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			// a real error happened! you should change your function return
// 			// to "(bool, error)" and return "false, err" here
// 			log.Print(err)
// 		}

// 		return false
// 	}

// 	return true
// }

//scan through database for name
// if {
// 	name is in the database
//	var account = true
// } else {
//	var account = false
//}

// kinda working
// 	var account = false
// 	return account
// }

// func d() bool {
//     var e bool
//     return e
// }

// if d() {
//     fmt.Printf("true")
// }

# Notes and unused code from main

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

	// var tolocation string
		// fmt.Println("Sorry not an option please choose [S]hop [A]rena [B]ank")
		// fmt.Scanln(&tolocation)
		// var tolocationlower = strings.ToLower(tolocation)