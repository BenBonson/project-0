package name

import (
	"fmt"
)

//GetName gets player name and returns it
func GetName() string {
	//get player name
	var Name string
	fmt.Println("Hello adventurer please state your name.")
	fmt.Scanln(&Name)
	return Name
}
