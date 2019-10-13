package shop

import (
	"fmt"
	"strings"

	"github.com/190930-UTA-CW-Go/project-0/cfunds"
)

//Shop for when the player is at the shop
func Shop() {
	fmt.Println("Welcome to the Shop. See anything you like?")
	cfunds := cfunds.Cfunds()
	fmt.Println("Currently holding", cfunds, "Gold")
	//array with options
	var inventory [7]string
	inventory[0] = "[P]late Mail 30g \n" //add descritpions once they do stuff
	inventory[1] = "[S]hort Sword 10g \n"
	inventory[2] = "[L]ong Sword 30g \n"
	inventory[3] = "[G]auntlets 15g \n"
	inventory[4] = "[B]oots 5g \n"
	inventory[5] = "[C]arrot of Death 300g \n"
	inventory[6] = "[E]xit shop \n"
	fmt.Println(inventory)
	var moneyspend string
	fmt.Scanln(&moneyspend)
	var moneyspendlower = strings.ToLower(moneyspend)
	//need to Check if they can afford it and subtract currently held funds
	if moneyspendlower == "p" {
		fmt.Println("Aquired Plate Mail")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "s" {
		fmt.Println("Aquired Short sword")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "s" {
		fmt.Println("Aquired Long sword")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "g" {
		fmt.Println("Aquired Gauntlets")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "b" {
		fmt.Println("Aquired Boots")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "c" {
		fmt.Println("Aquired Carrot of Death")
		fmt.Println("Thanks for the purchase now get out so I can spend my new gold.")
	} else if moneyspendlower == "e" {
		fmt.Println("Thats too bad.")
	} else {
		fmt.Println("Clearly you must have hit your head in the arena \ncause I have no idea what your going on about just get out and relearn how to speak.")
	}
}
