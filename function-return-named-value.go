package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {

	firstName = "tegar"
	middleName = "pratama"
	lastName = "p00"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
