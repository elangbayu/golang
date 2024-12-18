package main

import "fmt"

func main() {
	nilai32 := 32768
	nilai64 := int64(nilai32)
	nilai16 := int16(nilai32) // Akan jadi negatif (unexpected value) karena nilai32 melebihi kapasitas dari int16 (Number Overflow)

	fmt.Println(nilai32, nilai64, nilai16)

	convertedString := string("Elang"[2])
	fmt.Println(convertedString)
}
