package main

import "fmt"

func main() {
	var name string
	name = "Elang"
	fmt.Println(name)
	name = "Bayu"
	fmt.Println(name)

	// Jika langsung diinisiasi valuenya, maka tidak wajib ngasih type
	var nameDirect = "Elang"
	fmt.Println(nameDirect)
	nameDirect = "Bayu"
	fmt.Println(nameDirect)

	// Dapat disederhanakan lagi untuk variable yang langsung diinisasi valuenya
	nameSimple := "Elang"
	fmt.Println(nameSimple)
	nameSimple = "Bayu"
	fmt.Println(nameSimple)

	// Declare multiple variable
	var (
		firstName = "Elang"
		middleName = "Bayu"
	)
	fmt.Println(firstName, middleName)
}
