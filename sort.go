package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Tegar", 20},
		{"Bani", 19},
		{"Sigit", 17},
		{"Maulana", 18},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
