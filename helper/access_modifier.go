package helper

/*
Agar sebuah variable atau function bisa diakses dari luar package, maka huruf pertama harus kapital
Jika menggunakan huruf kecil, maka hanya bisa diakses dari package yang sama
*/

var version = "1.0.0" // Tidak bisa diakses dari luar package helper
var Application = "Go" // Bisa diakses dari luar package


// Tidak bisa diakses dari luar package helper
func sayGoodBye(name string) string {
	return "Goodbye, " + name
}
