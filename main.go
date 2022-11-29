package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var mf person
	fmt.Printf("%+v", mf)

	mf.firstName = "King"
	mf.lastName = "Geedorah"
	fmt.Println(mf.firstName)
	fmt.Println(mf.lastName)
}
