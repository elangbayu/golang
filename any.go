package main

import "fmt"

func Ups() interface{} {
	return 1
}

// interface{} adalah interface kosong, bisa juga diganti dengan keyword any

func Ups2() any {
	return "Bisa return tipe data apapun"
}

func main() {
	var kosong interface{} = Ups()
	var kosong2 any = Ups2()
	fmt.Println(kosong)
	fmt.Println(kosong2)
}
