package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	mf := person{
		firstName: "MF",
		lastName:  "DOOM",
		contactInfo: contactInfo{
			email:   "MF@DOOM",
			zipCode: 33333,
		},
	}

	mf.updateName("Metal Face")
	mf.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}
