package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Elang",
		"address": "Yogya",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	// Jika mengakses key yang tidak ada, maka akan return default value
	fmt.Println(person["last_name"])

	book := map[string]string{
		"author": "Elang",
		"title":  "Golang",
		"ups":    "salah",
	}

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)
}
