package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Vicry", "Fahreza", "Jean"})
	_ = writer.Write([]string{"Rian", "Miko", "Vera"})
	_ = writer.Write([]string{"Lucia", "Alice", "Moro"})

	writer.Flush()
}
