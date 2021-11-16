package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Tegar",
		"address": "Medan",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "tgr"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
