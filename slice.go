package main

import "fmt"

func main() {
	names := [...]string{"Elang", "Bayu", "Segara", "Wiwit", "Widowati", "Segara"}
	// Ambil mulai index ke-4 sampai index ke 6-1 (5)
	slice := names[4:6]
	fmt.Println(slice)
	// Ambil mulai index pertama sampai ke-3
	slice1 := names[:3]
	fmt.Println(slice1)
	// Ambil mulai index ke-3 sampai terakhir
	slice2 := names[3:]
	fmt.Println(slice2)
	// Ambil semua index
	slice3 := names[:]
	fmt.Println(slice3)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daysSlice := days[5:]
	fmt.Println(daysSlice)
	// Saat ubah slice yg dibuat dari array, maka data asli dari array akan ikut berubah
	daysSlice[0] = "Sabtu Baru"
	daysSlice[1] = "Minggu Baru"
	fmt.Println(daysSlice)
	fmt.Println(days)

	// Append akan membuat array baru dari array yang akan ditambahkan
	daysSlice1 := append(daysSlice, "Libur Baru")
	daysSlice1[0] = "Sabtu Lama" // Mengubah index pertama dari daySlice1
	fmt.Println(daysSlice) // Tidak berubah
	fmt.Println(daysSlice1)
	fmt.Println(days) // Tidak berubah

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Elang"
	newSlice[1] = "Elang"
	// newSlice[2] = "Elang" // Error, harus gunakan append
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	// Karena capacity newSlice adalah 5, dan ini hanya append 1, maka newSlice2 akan tetap refer ke newSlice
	newSlice2 := append(newSlice, "Elang")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Bayu"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // index pertama dari newSlice akan berubah jadi Bayu karena newSlice2 refer ke newSlice

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice, toSlice)

	// Perbedaan inisiasi array dan slice
	iniArray := [...]int{1,2,3,4,5}
	iniJugaArray := [3]int{1,2,3}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniJugaArray)
	fmt.Println(iniSlice)
}
