package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	tgr := Man{"tegar"}
	tgr.Married()

	fmt.Println(tgr.Name)
}
