package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Local())
	utc := time.Date(2006, time.January, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	// RFC3339 format waktu
	formatter := "2006-01-02 15:04:05"
	// value := "2009-02-04 17:10:10"
	value := "SALAH"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
}
