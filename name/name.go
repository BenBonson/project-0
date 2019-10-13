package name

import (
	"fmt"
)

//Name is player name
var Name string

//GetName gets player name and returns it
func GetName() string {
	//get player name
	//var Name string
	if len(Name) > 0 {
		//check to see if it runs
		//fmt.Println("Sup")
	} else {
		fmt.Println("Hello adventurer please state your name.")
		fmt.Scanln(&Name)
	}
	//playername := Name
	return Name
}

//Name prints player name
// func Name() {
// 	fmt.Println(playername)
// }

// //GetName gets player name and returns it
// func GetName() string {
// 	//get player name
// 	var Name string
// 	// if len(Name) > 0 {
// 	// 	fmt.Println("Sup")
// 	// } else {
// 	fmt.Println("Hello adventurer please state your name.")
// 	fmt.Scanln(&Name)
// 	//}
// 	return Name
// }

// // func name() {
// // 	name := GetName(Name)
// // }

// func name() {
// 	name := GetName(Name)
// }

//improves name getting
// reader := bufio.newReader(os.Stdin)
//     var name string
//     fmt.Println("What is your name?")
//     name, _ := reader.readString("\n")
