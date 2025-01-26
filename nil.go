package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

/*
nil hanya bisa digunakan untuk return type berikut: interface, function, map, slice, pointer, dan channel
selain type tersebut, nil tidak bisa digunakan sebagai return type
berikut adalah contoh function dengan return type string, tidak bisa return type nil

	func ContohError(name string) string {
		if name == "" {
			return nil //Error karena tipe data string tidak bisa return nil
		} else {
			return name
		}
	}
*/

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data map nil")
	} else {
		fmt.Println(data["name"])
	}
}
