package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	/* Jika tipe tidak sesuai, maka akan panic
	var resultInt int = result.(int)
	fmt.Println(resultInt)
	*/

	// For safety, kita bisa melakukan pengecekan tipe data sebelum conversions
	var result2 any = random()
	switch result2.(type) {
	case string:
		// tidak perlu conversions lagi karena sudah masuk dari case
		fmt.Println("String", result2)
	case int:
		fmt.Println("Int", result2)
	default:
		fmt.Println("Unknown type", result2)
	}
}
