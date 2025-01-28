package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local()) // Local TimeZone
	utc := time.Date(1998, time.November, 14, 15, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"
	value := "2020-10-10 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err == nil {
		fmt.Println(valueTime)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Year())
}
