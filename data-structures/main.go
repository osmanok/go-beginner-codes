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
	gok := person{
		firstName: "Gok",
		lastName:  "Okuyan",
		contactInfo: contactInfo{
			email:   "gok@exp.com",
			zipCode: 12345,
		},
	}

	gok.updateName("Gokhan")

	gok.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
