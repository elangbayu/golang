package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func filter(name string) string {
	if name == "Anjing" {
		return "***"
	}
	return name
}

// Agar rapi, penulisan kontrak function as parameter bisa diganti dengan type declarations

type SpamFilter func(string) string

func sayHelloWithSpamFilter(name string, filter SpamFilter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func main() {
	sayHelloWithFilter("Elang", filter)
	// as value
	filters := filter
	sayHelloWithFilter("Anjing", filters)

	sayHelloWithSpamFilter("Anjing", filters)
}
