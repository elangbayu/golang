package main

import (
	"fmt"
	"os"
)

func main() {
	argument := os.Args
	for _, arg := range argument {
		fmt.Println(arg)
	}

	host, err := os.Hostname()
	if err == nil {
		fmt.Println(host)
	} else {
		fmt.Println(err.Error())
	}
}
