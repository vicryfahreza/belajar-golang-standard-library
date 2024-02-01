package main

import (
	"fmt"
	"math"
)

func main() {
	ceilNumber := math.Ceil(2.7)
	roundNumber := math.Round(2.4)
	floorNumber := math.Floor(2.4)
	fmt.Println(ceilNumber)
	fmt.Println(roundNumber)
	fmt.Println(floorNumber)
	fmt.Println(math.Max(10, 11))
	fmt.Println(math.Min(10, 11))
}
