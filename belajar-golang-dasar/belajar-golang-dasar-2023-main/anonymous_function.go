package main

import "fmt"

// type Blacklist func(string) bool

// func registerUser(name string, blacklist Blacklist) {
// 	if blacklist(name) {
// 		fmt.Println("You are blocked", name)
// 	} else {
// 		fmt.Println("Welcome", name)
// 	}
// }

// func main() {
// 	blacklist := func(name string) bool {
// 		return name == "anjing"
// 	}
// 	registerUser("eko", blacklist)

// 	registerUser("anjing", func(name string) bool {
// 		return name == "anjing"
// 	})
// }

type whiteList func(string) bool

func acceptUser(name string, whitelist whiteList)  {

if whitelist(name) {
	fmt.Println("you are accepted", name)

} else {
	fmt.Println("you are blocked", name)
}
}

func main() {
whitelist := func (name string) bool {
	if name == "anjing" {
		return true
	} 

	return false
}
acceptUser("anjing", whitelist)
}
