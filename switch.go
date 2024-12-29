package main

import "fmt"

func main() {
	name := "Elang"

	switch name {
	case "Elang":
		fmt.Println("Hello Elang")
	case "Bayu":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Selamat datang")
	}

	// switch short-statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama aman", length)
	default:
		fmt.Println("Nama terlalu panjang", length)
	}

	//switch tanpa kondisi. basically if-else statement
	switch {
	case name == "Elang":
		fmt.Println("Hello Elang")
	case name == "Budi":
		fmt.Println("Hello Elang")
	default:
		fmt.Println("Selamat datang")
	}
}
