package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	result2, err := strconv.Atoi("100")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result2)
	}

	binary := strconv.FormatInt(1998, 2)
	fmt.Println(binary)

	strInt := strconv.Itoa(1998)
	fmt.Println(strInt)
}
