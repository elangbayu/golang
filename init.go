package main

import (
	"fmt"
	"learn-golang/database"
	_ "learn-golang/internal" // This will execute the init method within internal without the need to execute any of its method
)

func main() {
	fmt.Println(database.GetDatabase())
}
