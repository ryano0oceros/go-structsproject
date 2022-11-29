package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	mf := person{
		firstName: "MF",
		lastName:  "DOOM",
		contact: contactInfo{
			email:   "MF@DOOM",
			zipCode: 33333,
		},
	}
	fmt.Println(mf.firstName)
	fmt.Println(mf.lastName)
	fmt.Println(mf.contact)
}
