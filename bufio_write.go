package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello saya Nami\n")
	_, _ = writer.WriteString("Saya suka berlayar\n")

	writer.Flush()
}
