package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func ChangeAddress(address Address) {
	address.Country = "Indonesia"
}

func ChangeAddressWithPointer(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeAddress(address)
	/*
	Although the address supposed to be changed within the ChangeAddress func, it will not.
	Simply because the address that changed within the ChangeAddress function is different from the one that printed
	This is because we pass it by value, thus it's just copy new address and assign to it
	Below, we print the initial address that's not changed at all
	*/
	fmt.Println(address)

	// Instead, use pointer to access the same address
	ChangeAddressWithPointer(&address)
	fmt.Println(address)
}
