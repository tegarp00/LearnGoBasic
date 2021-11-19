package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	fmt.Println(arg)
	//fmt.Println(arg[1])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error :", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
