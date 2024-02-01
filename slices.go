package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Johny", "Meni", "Rika", "Goro"}
	num := []int{10, 10, 5, 25, 30}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(num))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(num))
	fmt.Println(slices.Contains(names, "Johny"))
	fmt.Println(slices.Index(names, "Johny"))
	fmt.Println(slices.Index(names, "Rika"))
}
