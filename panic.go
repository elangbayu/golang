package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Recovered", message)
}

func endAppWithRecover() {
	fmt.Println("End app")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Ups error")
	}
}

func runAppWithRecover(err bool) {
	defer endAppWithRecover()
	if err {
		panic("Ups error")
	}
}

func main() {
	runApp(true)
	runAppWithRecover(true)
}
