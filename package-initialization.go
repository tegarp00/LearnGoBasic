package main

import (
	"fmt"
	"learngo/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
