package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	for _, args := range arg {
		fmt.Println(args)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
	// command untuk menambahkan argument satu persatu
	// go run --nama file go-- vicry fahreza hijack
	// langsung
	// go run --nama file go-- "vicry fahreza hijack"

}
