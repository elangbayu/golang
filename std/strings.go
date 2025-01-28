package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Elang Segara", "gara"))
	fmt.Println(strings.Split("Elang Segara", " "))
	fmt.Println(strings.ToLower("Elang Segara"))
	fmt.Println(strings.ToUpper("Elang Segara"))
	fmt.Println(strings.Trim(" Elang Segara ", " "))
	fmt.Println(strings.ReplaceAll("Elang Segara", "Segara", "Bayu"))
}
