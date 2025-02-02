package helper

import (
	"fmt"
	"testing"
)

/*
Secara default, go test hanya akan running semua test yang ada di current folder/package
Jika ingin run dari root folder, lakukan seperti ini "go test ./.."
*/

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Elang")
	if result != "Hello Elang" {
		t.Fail()
	}

	fmt.Println("Kode ini akan dijalankan karena hanya Fail()")
}

func TestHelloWorldNow(t *testing.T) {
	result := HelloWorld("Elang")
	if result != "Failed Elang" {
		t.FailNow()
	}

	fmt.Println("Kode ini tidak akan dijalankan karena menggunakan FailNow()")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Elang")
	if result != "Failed Elang" {
		t.Error("Ini sama dengan Fail() hanya saja ada pesan errornya")
	}

	fmt.Println("Kode ini akan dijalankan karena menggunakan Error()")
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Elang")
	if result != "Failed Elang" {
		t.Fatal("Ini sama dengan FailNow() hanya saja ada pesan errornya")
	}

	fmt.Println("Kode ini tidak akan dijalankan karena menggunakan Fatal()")
}
