package main

import "fmt"

func getFullName(firstName string, lastName string) (string, string) {
	return "First name: " + firstName, "Last name: " + lastName
}

func main() {
	// Ambil semua return value
	fmt.Println(getFullName("Elang", "Segara"))

	// Ambil salah satu return value
	firstName, _ := getFullName("Elang", "Segara")
	fmt.Println("Firstname only: ", firstName)
}
