package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "tegar"
	names[1] = "pratama"
	names[2] = "senju"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...]int{90, 80, 70, 11}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println(len(names))
	fmt.Println(len(values))
}
