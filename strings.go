package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Vicry Fahreza", "Vicry"))
	fmt.Println(strings.Split("Vicry Fahreza", ""))
	fmt.Println(strings.ToLower("Vicry Fahreza"))
	fmt.Println(strings.ToUpper("Vicry Fahreza"))
	fmt.Println(strings.Trim("   Vicry Fahreza   ", " "))
	fmt.Println(strings.ReplaceAll("Vicry Fahreza Vicry Maharja", "Vicry", "Seno"))
}
