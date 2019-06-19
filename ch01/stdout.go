package main

import (
	"io"
	"os"
)

func main() {
	incomming := "Plase give me agrument"
	if len(os.Args) > 1 {
		incomming = os.Args[1]
	}
	io.WriteString(os.Stdout, incomming)
	io.WriteString(os.Stdout, "\n")
}
