package main

import "fmt"

func main() {
	var (
		a = 1
		b = 2
	)
	c := a + b
	fmt.Println(c)

	// Augmented assignments
	i := 10
	fmt.Println(i)
	i += 5
	fmt.Println(i)

	// Unary operator
	j := 1
	fmt.Println(j)
	j++
	fmt.Println(j)
	j--
	fmt.Println(j)
}
