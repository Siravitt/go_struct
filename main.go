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

type employees []person

func main() {
	jim := createPerson("Jim", "Haland", "jim@gmail.com", 10200)
	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func createPerson(firstName string, lastName string, email string, zipCode int) person {
	return person{
		firstName: firstName,
		lastName:  lastName,
		contactInfo: contactInfo{
			email,
			zipCode,
		},
	}
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
