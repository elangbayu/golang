package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		fmt.Println("Name ", valueType.Field(i).Name)
		fmt.Println("Type ", valueType.Field(i).Type)
		fmt.Println("Type ", valueType.Field(i).Tag.Get("required"))
		fmt.Println("Type ", valueType.Field(i).Tag.Get("max"))
	}
}

func isValid(data any) (result bool) {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{"Elang"})
	readField(Person{"Segara", "jogja", "email"})
	person := Person{"Elang", "jogja", ""}
	fmt.Println(isValid(person))
}
