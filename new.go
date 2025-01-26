package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := new(Address) // This will create an empty pointer
	address2 := address1
	address2.City = "Sidoarjo"
	fmt.Println(address1)
	fmt.Println(address2)

	// Or you can define empty pointer like this
	address3 := &Address{}
	address4 := address3
	address4.City = "Yogyakarta"
	fmt.Println(address3)
	fmt.Println(address4)
}
