package main

import "fmt"

type Man struct {
	Name string
}

func (man Man) Married() {
	man.Name = "Mr. " + man.Name
}

func (man *Man) MarriedPointer() {
	man.Name = "Mr. " + man.Name
}

func main() {
	elang := Man{"Elang"}
	elang.Married()

	/*
	We expect the name of elang changed to "Mr. Elang", but turns out it isn't
	This is because when we execute elang.Married(), it will copy the value of elang to new pointer within the function
	So, what changed inside the function doesn't impact the actual elang that we print
	*/
	fmt.Println(elang.Name)

	// Instead, use asterisk to actually change the value
	elang.MarriedPointer()
	fmt.Println(elang.Name)
}
