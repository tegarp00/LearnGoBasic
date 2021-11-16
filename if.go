package main

import "fmt"

func main() {

	var name = "tegar"

	if name == "tegfar" {
		fmt.Println("Hello gar...")
	} else {
		fmt.Println("ydhlah")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah pas")
	}
}
