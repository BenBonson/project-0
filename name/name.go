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
		fmt.Println("Hello adventurer please state your name no spaces.")
		fmt.Scanln(&Name)
		// in := bufio.NewReader(os.Stdin)
		// Name, _ := in.ReadString('\n')
	}
	return Name
}

//to do
//get names with spaces

//Attempts to get first and last name
// package name

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// //Name is player name
// var Name string

// //GetName gets player name and returns it
// func GetName() string {
// 	//get player name
// 	//var Name string
// 	if len(Name) > 0 {
// 		//check to see if it runs
// 		//fmt.Println("Sup")
// 	} else {
// 		fmt.Println("Hello adventurer please state your name.")
// 		// // String()
// 		// reader := bufio.NewReader(os.Stdin)
// 		scanner := bufio.NewScanner(os.Stdin)
// 		scanner.Scan() // use `for scanner.Scan()` to keep reading
// 		Name := scanner.Text()
// 		fmt.Println("captured:", &Name)
// 	}
// 	fmt.Println("WTF")
// 	return Name
// }

//Fname, err := strconv.Atoi(input)
// input, _ := reader.ReadString('\n')
// input = strings.TrimSuffix(input, "\n")

//String Reads user input past a space
// func String() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	bytes, _, _ := reader.ReadLine()
// 	fmt.Println(string(bytes))
// 	return string(bytes)
// }

// func String() {
// 	reader := bufio.NewReader(os.Stdin)
// 	return
// }
