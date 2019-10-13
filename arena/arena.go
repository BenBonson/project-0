package arena

import (
	"fmt"
	"strings"

	"github.com/190930-UTA-CW-Go/project-0/cfunds"
)

//Arena for when the player goes to the arena
func Arena() {
	fmt.Println("Sup wanna [f]ight or are you a [c]oward?")
	var fight string
	fmt.Scanln(&fight)
	var fightlower = strings.ToLower(fight)
	if fightlower == "f" {
		fmt.Println("Heck yah, so wanna bet? [Y] or [N]")
		cfunds := cfunds.Cfunds()
		fmt.Println("Currently holding", cfunds, "Gold")
		//option to bet
	} else {
		fmt.Println("Get lost then and stop wasting my time!")
	}
}

//array of fighters and stats

//acctual fights or random generator
