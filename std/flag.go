package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "root", "username database")
	password := flag.String("password", "root", "password database")
	host := flag.String("host", "localhost", "host database")
	port := flag.String("port", "1234", "port database")

	flag.Parse()

	fmt.Printf("Username: %s\nPassword: %s\nHost: %s\nPort: %s\n", *username, *password, *host, *port)
}
