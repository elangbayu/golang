package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // akan berhenti di angka 5
		}
		fmt.Println("Perulangan ke", i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // akan skip angka 5
		}
		fmt.Println("Perulangan ke", i)
	}
}
