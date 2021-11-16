package main

import "fmt"

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("Hello", nameFiltered)
// }

// func spamFilter(name string) string {
// 	if name == "Anjing" {
// 		return "..."
// 	} else if name == "Babi" {
// 		return "Kaulah"
// 	} else {
// 		return name
// 	}
// }

// Pakai type Declaration

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else if name == "Babi" {
		return "Kaulah"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("tegar", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	sayHelloWithFilter("Babi", filter)
}
