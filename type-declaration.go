package main

import "fmt"

func main() {
	type (
		NoKTP   string
		Married bool
	)

	var noKtpTgr NoKTP = "12798419481240000"
	var marriedStatus Married = false
	fmt.Println("No KTP :", noKtpTgr)
	fmt.Println("Status Menikah :", marriedStatus)
}
