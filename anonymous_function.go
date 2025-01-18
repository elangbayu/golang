package main

import "fmt"

type Blacklist func(string) bool

func register(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked")
	} else {
		fmt.Println("Welcome,", name)
	}
}

func main() {
	// Cara #1
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	register("Anjing", blacklist)

	// Cara #2
	register("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
