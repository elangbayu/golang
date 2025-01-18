package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 30))

	// jika ingin memasukkan slice sebagai param untuk variadic function, maka tambahkan ... sebagai suffix dari argumen yg dilempar
	sliceParam := []int{10, 20, 30}
	fmt.Println(sumAll(sliceParam...))
}
