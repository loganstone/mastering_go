package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	START = "!"
	MIN   = 0
	MAX   = 94
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	var LENGTH uint64 = 8

	switch len(os.Args) {
	case 2:
		n, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println("Bad incoming:", err)
			os.Exit(1)
		}
		LENGTH = n
	default:
		fmt.Println("Using default value:", LENGTH)
	}

	rand.Seed(time.Now().Unix())

	var i uint64 = 1
	for {
		n := random(MIN, MAX)
		// Helpful command
		// $ man ascii
		// 33 ~ 126
		c := string(START[0] + byte(n))
		fmt.Print(c)
		if i == LENGTH {
			break
		}
		i++
	}
	fmt.Println()
}
