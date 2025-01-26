package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	AnimalName string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (pet Animal) GetName() string {
	return pet.AnimalName
}

func main() {
	person := Person{Name: "Elang"}
	pet := Animal{AnimalName: "Bobi"}
	SayHello(person)
	SayHello(pet)
}
