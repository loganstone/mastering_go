package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, fmt.Sprintf("> %s\n", scanner.Text()))
	}
}
