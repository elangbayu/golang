package main

import "fmt"

func main() {
	name := "Elang"

	if name == "Elang" {
		fmt.Println("Hello Elang")
	} else if name == "Bayu" {
		fmt.Println("Hello Bayu")
	} else {
		fmt.Println("Selamat datang")
	}

	// if short-statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang", length)
	} else {
		fmt.Println("Nama aman", length)
	}
}
