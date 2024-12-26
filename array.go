package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Elang"
	names[1] = "Bayu"
	names[2] = "Segara"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Di luar kapasitas array
	// names[3] = "Error"

	// Langsung initiate value
	var values = [3]int {
		10,
		20,
		30,
	}

	fmt.Println(values)

	var defaultValuesInt = [3]int {
		10,
		20,
	}
	// Index terakhir adalah 0 (default value)
	fmt.Println(defaultValuesInt)

	var defaultValuesString = [3]string {
		"Elang",
		"Bayu",
	}
	// Index terakhir adalah empty string (default value)
	fmt.Println(defaultValuesString)

	// Menentukan ukuran array sebanyak data yg diinisiasi, hanya bisa dilakukan pada array yg diinisiasi langsung
	var predefinedArray = [...]int {
		10,
		20,
		30,
	}
	fmt.Println(len(predefinedArray))
}
