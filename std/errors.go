package main

import "fmt"
import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "elang" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("ValidationError")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("NotFoundError")
		} else {
			fmt.Println("Unknonw Error")
		}
	}
}
