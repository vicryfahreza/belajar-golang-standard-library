package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`a([a-z])i`)

	fmt.Println(regex.MatchString("adi"))
	fmt.Println(regex.MatchString("ari"))
	fmt.Println(regex.MatchString("aJi"))

	fmt.Println(regex.FindAllString("adi ari ami aJi a1u ali", 10))
}
