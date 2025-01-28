package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Elang", "Bayu", "Segara"}
	age := []int{10, 20, 30}
	fmt.Println(slices.Max(age))
	fmt.Println(slices.Min(age))
	fmt.Println(slices.Contains(names, "Bayu"))
	fmt.Println(slices.Index(names, "Segara"))
}
