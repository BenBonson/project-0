package main

import (
	"fmt"

	"github.com/190930-UTA-CW-Go/project-0/bank"

	"strings"
)

func main() {
	var name string
	fmt.Println("Hello adventurer please state your name.")
	fmt.Scanln(&name)

	if bank.Account() {
		fmt.Println("HOW")
	} else {
		fmt.Println("Welcome", name, "I see you do not have an account with us would you like to set one up? [Y] or [N]")
	}

	var acreate string
	fmt.Scanln(&acreate)
	var acreatelower = strings.ToLower(acreate)
	fmt.Println(acreatelower)
	if acreatelower == "y" {
		fmt.Println("Alright let me create that for you. Just so you know if you store your gold at the bank you will be unable to bet it however you wont lose it if you are beaten.")
	} else {
		fmt.Println("Well, good luck with that. Just so you know if you store your gold at the bank you will be unable to bet it however you wont lose it if you are beaten.")
	}

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
}
