package main

import "fmt"

func main() {

	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	for counter := 1; counter <= 11; counter++ {
		fmt.Println("Ke :", counter)
	}

	slice := []string{"tegar", "pratama", "p00"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// For range
	//names := []string{"tegar", "pratama", "p00"}
	for _, value := range slice {
		//fmt.Println("index", i, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "tegar"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
