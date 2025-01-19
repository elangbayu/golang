package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var cust Customer
	cust.Name = "Elang"
	cust.Address = "Sidoarjo"
	cust.Age = 24
	fmt.Println(cust)

	// Struct literals
	custNew := Customer{
		Name:    "Bayu",
		Address: "Yogyakarta",
		Age:     25,
	}
	fmt.Println(custNew)

	// Struct literals #2
	custNew2 := Customer{"Segara", "Surabaya", 26}
	fmt.Println(custNew2)
}
