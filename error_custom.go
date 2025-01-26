package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Id empty"}
	}

	if id != "elang" {
		return &notFoundError{"data not found"}
	}

	//ok
	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error", validationError.Error())
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error", notFoundError.Error())
		} else {
			fmt.Println("Unknown error", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}

	err2 := SaveData("bayu", nil)
	if err2 != nil {
		// Or use switch
		switch finalError := err2.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error", finalError.Error())
		default:
			fmt.Println("unknown error", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
