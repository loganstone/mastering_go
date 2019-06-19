package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const startIdx = 2
const bitSize = 64

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("Please give one or more floats.")
	}

	min, _ := strconv.ParseFloat(os.Args[1], bitSize)
	max, _ := strconv.ParseFloat(os.Args[1], bitSize)

	for i := startIdx; i < len(os.Args); i++ {
		n, _ := strconv.ParseFloat(os.Args[i], bitSize)
		if n < min {
			min = n
		} else if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min, "Max:", max)
}
