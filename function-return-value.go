package main

import "fmt"

func getName(name string) string {
	if name == "" {
		return "Hello Broo..."
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(getName("p00"))
}
