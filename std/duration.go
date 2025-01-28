package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := 100 * time.Second
	duration2 := 10 * time.Millisecond
	duration3 := duration1 - duration2

	fmt.Println(duration3) // Print in human readable format
	fmt.Printf("duration: %d\n", duration3) // Print in int format
}
