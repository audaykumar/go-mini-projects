package main

import (
	"fmt"
)

type contactInfo struct {
	email    string
	phoneNum string
}

type person struct {
	firstname string
	lastname  string
	// contact   contactInfo
	contactInfo
}

func main() {

	// obj := person{"Aniket", "Udaykumar"}

	// obj := person{firstname: "Aniket", lastname: "Udaykumar"}

	// var obj person
	// obj.firstname = "Aniket"
	// obj.lastname = "Udaykumar"

	// fmt.Println(obj)

	obj := person{
		firstname: "Aniket",
		lastname:  "Udaykumar",
		// contact: contactInfo{
		// 	email:    "auk.xyz.com",
		// 	phoneNum: "987654321",
		// },
		contactInfo: contactInfo{
			email:    "auk.xyz.com",
			phoneNum: "987654321",
		},
	}

	obj.print()
	obj.printFirstName()
	obj.updateName("XYZ")
	obj.printFirstName()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) printFirstName() {
	fmt.Printf("%+v\n", p.firstname)
}

func (p person) updateName(newFirstName string) {
	p.firstname = newFirstName
}
