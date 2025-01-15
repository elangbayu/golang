package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(getHello("Elang"))
}
