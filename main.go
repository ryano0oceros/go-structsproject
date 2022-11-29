package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	mf := person{firstName: "MF", lastName: "DOOM"}
	fmt.Println(mf)
	fmt.Println(mf.firstName)
	fmt.Println(mf.lastName)
}
