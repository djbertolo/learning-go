package main

import "fmt"

type Contact struct {
	email       string
	phoneNumber string
}

func main() {
	contacts := make(map[string]Contact, 0)
	getContactInfo := func(name string) {
		greeting := fmt.Sprintf("%v's contact info:\n", name)
		info, ok := contacts[name]
		if !ok {
			greeting += "Contact not found"
		} else {
			greeting += fmt.Sprintf("email: %v\nnumber: %v", info.email, info.phoneNumber)
		}

		fmt.Println(greeting)
	}

	contacts["Alex"] = Contact{
		email:       "alex@gmail.com",
		phoneNumber: "(123)456-7890",
	}

	contacts["Steve"] = Contact{
		email:       "steve@gmail.com",
		phoneNumber: "(987)654-3210",
	}

	getContactInfo("Alex")
	getContactInfo("Steve")
	getContactInfo("Mike")
}
