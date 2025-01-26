package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Wonosari", "DIY", "Indonesia"}
	address3 := &address1 // This will refer to address1, thus what happens with address3 will impact address1
	address3.City = "Mojokerto"
	fmt.Println(address1)
	fmt.Println(address3)

	// Asterisk is used to force all that refer to the same pointer to move to the new pointer
	*address3 = Address{"Sidoarjo", "Jawa Timur", "Indonesia"} // This will make the address1 to also move to this new pointer
	fmt.Println(address1)
	fmt.Println(address3)
}
