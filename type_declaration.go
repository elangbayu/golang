package main

import "fmt"

func main() {
	type NoKTP string
	var ktpElang NoKTP = "123456789"
	contohString := "22222222222"
	contohKtp := NoKTP(contohString)
	fmt.Println(ktpElang, contohKtp)
}
