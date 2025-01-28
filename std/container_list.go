package main

import (
	"container/list"
	"fmt"
)

/*
Container list adalah implementasi dari double linked list
*/

func main() {
	data := list.New()
	data.PushBack("Elang")
	data.PushBack("Bayu")
	data.PushBack("Segara")

	head := data.Front() // Dapat data paling depan (pertama kali diinsert)
	fmt.Println(head.Value)

	headNext := head.Next() // Dapat data setelah head
	fmt.Println(headNext.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
