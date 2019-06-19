package main

import (
	"io"
	"os"
)

func main() {
	incomming := "Plase give me agrument."
	if len(os.Args) > 1 {
		incomming = os.Args[1]
	}
	io.WriteString(os.Stdout, "This is Standart output.\n")

	// stderr redirection to file
	// $ go run stderr.go 2>somewhere

	// ignore stderr
	// $ go run stderr.go 2>/dev/null

	// output redirection to file and stderr redirection to stdout
	// $ go run stderr.go >somewhere 2>&1

	// ignore all output
	// $ go run stderr.go >/dev/null 2>&1
	io.WriteString(os.Stderr, incomming)
	io.WriteString(os.Stderr, "\n")
}
