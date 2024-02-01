package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/hello.go"))
	fmt.Println(path.Base("hello/hello.go"))
	fmt.Println(path.Ext("hello/hello.go"))
	fmt.Println(path.Join("hello", "world", "main.go"))
}
