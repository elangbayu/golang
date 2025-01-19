package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (cust Customer) sayHello(name string) {
	fmt.Println("Hello", name)
	fmt.Println("Your data", cust)
}

func main() {
	elang := Customer{"Elang", "Sidoarjo", 24}
	elang.sayHello(elang.Name)
}
