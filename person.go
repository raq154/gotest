package main

import "fmt"

type ContactInfo struct {
	email   string
	zipcode string
}

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (p Person) print() {
	fmt.Printf("%s %s %s %s \n", p.firstName, p.lastName, p.contactInfo.email, p.contactInfo.zipcode)
}
func (p *Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
