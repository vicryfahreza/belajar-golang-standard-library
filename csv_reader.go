package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "vicry,fahreza,fuhua\n" +
		"Rian,Wanto,Shopia\n" +
		"Mega,Wija,Teri\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

}
