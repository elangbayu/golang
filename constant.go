package main

func main() {
	// Jika constant tidak digunakan tidak masalah
	const firstName = "Elang"

	// error karena constant tidak boleh diganti valuenya
	// firstName = "Bayu"

	// Multiple constant
	const (
		middleName = "Bayu"
		lastName = "Segara"
	)
}
