package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

func main() {
	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "jim@jim.com",
	// 		zipCode: 100,
	// 	},
	// }
	// fmt.Println(jim)
	// fmt.Printf("%+v", jim)

	kevin := person{
		firstName: "qiongjie",
		lastName:  "zhang",
		contactInfo: contactInfo{
			email:   "qiongjie@zhang.com",
			zipCode: 99,
		},
	}
	kevin.updateName("boyan")
	kevin.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
