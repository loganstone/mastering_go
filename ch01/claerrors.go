package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Please give one or more floats.")
	}

	initIdx := 1
	startIdx := 2
	bitSize := 64
	err := errors.New("an error failed min, max initialization")

	var n float64

	for err != nil {
		if initIdx >= len(os.Args) {
			log.Fatalln("None of the arguments is a float")
		}
		n, err = strconv.ParseFloat(os.Args[initIdx], bitSize)
		initIdx++
	}

	min, max := n, n

	for i := startIdx; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], bitSize)
		if err == nil {
			if n < min {
				min = n
			} else if n > max {
				max = n
			}
		}
	}

	fmt.Println("Min:", min, "Max:", max)
}
