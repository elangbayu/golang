package main

import "fmt"

func getName(firstName string, lastName string) (firstNameValue string, lastNameValue string) { // can be (firstName, lastName string) since they are the same type
	firstNameValue = firstName
	lastNameValue = lastName
	return // this will automatically return all the return values. Only works with named return since it's already assigned with its default value
}

func main() {
	firstName, lastName := getName("Elang", "Segara")
	fmt.Println("Full name:", firstName, lastName)
}
