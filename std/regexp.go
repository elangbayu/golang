package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e[a-z]+`)
	fmt.Println(regex.MatchString("elang"))
	fmt.Println(regex.MatchString("segara"))
	fmt.Println(regex.MatchString("Bayu"))
	fmt.Println(regex.MatchString("seg4ra"))
	fmt.Println(regex.FindAllString("elang bayu segara", 2))
}
