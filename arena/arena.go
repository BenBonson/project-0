package arena

import (
	"fmt"
	"strings"
)

//Arena for when the player goes to the arena
func Arena() {
	fmt.Println("Sup wanna [f]ight or are you a [c]oward?")
	var fight string
	fmt.Scanln(&fight)
	var fightlower = strings.ToLower(fight)
	if fightlower == "f" {
		fmt.Println("Heck yah, so wanna bet? [Y] or [N]")
	} else {
		fmt.Println("Get lost then and stop wasting my time!")
	}
}
