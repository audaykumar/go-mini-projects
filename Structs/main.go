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
	// Creating a pointer object
	// objPointer := &obj
	// objPointer.updateName("XYZ")

	// Shortcut to using pointer.
	// We can directly use the type of struct while sending to the pointer
	obj.updateName("QWE")
	obj.printFirstName()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) printFirstName() {
	fmt.Printf("%+v\n", p.firstname)
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstname = newFirstName
}
