package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Pass by value (the default)
	address1 := Address{"Wonosari", "DIY", "Indonesia"}
	address2 := address1 // This will copy the address1 to address2, thus what happens with address2 will not impact address1
	address2.City = "Sidoarjo"
	fmt.Println(address1)
	fmt.Println(address2)

	// Pass by reference
	address3 := &address1 // This will refer to address1, thus what happens with address3 will impact address1
	address3.City = "Mojokerto"
	fmt.Println(address1)
	fmt.Println(address3)
}
