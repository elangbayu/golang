package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
Container ring adalah implementasi dari circular linked list
*/

func main() {
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// Do adalah seperti map di javascript
	data.Do(func(value any) {
		fmt.Println(value)
	})
}
