package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Medan", "Sumut", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	var address3 *Address = new(Address)

	address3.City = "tj.anom"
	fmt.Println(address3)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	alamatChange := &alamat
	ChangeCountryToIndonesia(alamatChange)
	fmt.Println(alamat)
}
