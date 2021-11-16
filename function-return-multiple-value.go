package main

import "fmt"

func getFullName() (string, string) {
	return "tgr", "p00"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
