package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Data Not Found")
)

func getById(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Vicry" {
		return NotFoundError
	}
	return nil
}

func main() {
	err := getById("Vicry")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error : ", err.Error())
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error : ", err.Error())
		} else {
			fmt.Println("Unknown Error : ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
