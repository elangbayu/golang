package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai")

	// for with statement
	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke ", count)
	}

	fmt.Println("Selesai")

	names := []string{"Elang", "Bayu", "Segara"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// range
	// karena names adalah slice/array, maka return dari range adalah index dan valuenya
	// andaikan names adalah map, maka return dari range adalah key dan valuenya
	// jika return value tidak dipakai, maka variable dapat diganti dengan _
	// contoh: for _, name := range names hanya akan dapat mengakses value dari array names, tidak dengan indexnya
	for index, name := range names {
		fmt.Println("Index", index, "name", name)
	}
}
