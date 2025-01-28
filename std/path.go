package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	pathh := "hello/world.go"
	fmt.Println(path.Dir(pathh))  // hello
	fmt.Println(path.Base(pathh)) // world.go
	fmt.Println(path.Ext(pathh))  // .go
	fmt.Println(path.Join("hello", "world", "main.go"))

	// Filepath
	fmt.Println(filepath.Dir(pathh))
	fmt.Println(filepath.Base(pathh))
	fmt.Println(filepath.Ext(pathh))
	fmt.Println(filepath.IsAbs(pathh))   // Apakah mulai dari root?
	fmt.Println(filepath.IsLocal(pathh)) // Apakah relative path?
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
