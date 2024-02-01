package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
	// command untuk menambahkan argument melalui prompt
	// go run flag.go -username=vicry -password=rahasia -host=123.231.23.3 -port=5505
	// go run flag.go -username=vicry -password="rahasisa banget" -host=123.231.23.3 -port=5505
}
