package main

import "fmt"

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	say := sayHello
	fmt.Println(say("Elang"))
}
