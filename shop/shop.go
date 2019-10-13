package shop

import "fmt"

//Shop for when the player is at the shop
func Shop() {
	fmt.Println("Welcome to the Shop. See anything you like?")
	//array with options
	var inventory [6]string
	inventory[0] = "[P]late Mail 30g \n"
	inventory[1] = "[S]hort Sword 10g \n"
	inventory[2] = "[L]ong Sword 30g \n"
	inventory[3] = "[G]auntlets 15g \n"
	inventory[4] = "[B]oots 5g \n"
	inventory[5] = "[C]arrot of death 300g \n"
	fmt.Println(inventory)
}

//add ability to actualy buy stuff
