package main

import "fmt"

func main() {

	name := "tgr"

	switch name {
	case "tgr":
		fmt.Println("Hello tegar")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("ydhlah...")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
