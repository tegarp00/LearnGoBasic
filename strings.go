package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Tegar Pratama", "Tegar"))

	fmt.Println(strings.Split("Tegar Pratama", " "))

	fmt.Println(strings.ToLower("tegar pratama"))
	fmt.Println(strings.ToUpper("tegar pratama"))

	fmt.Println(strings.Trim("   Tegar Pratama       ", " "))

	fmt.Println(strings.ReplaceAll("Tegar pra p00", "pra", "pratama"))

}
