package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging() // Make sure logging is called even this method is error
	fmt.Println("Running...")
}

func main() {

}
