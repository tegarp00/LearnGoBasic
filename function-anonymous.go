package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "Admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "Root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin" || name == "root"
	}

	registerUser("admin", blacklist)
	registerUser("tgr", blacklist)
	registerUser("root", blacklist)
	// registerUser("root", func(name string) bool {
	// 	return name == "root"
	// })
}
