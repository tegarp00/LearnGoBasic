package main

import "fmt"

func sayHello() {
	fmt.Println("Hello Bro...")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHello()
	sayHelloTo("tgr", "p00")
}
