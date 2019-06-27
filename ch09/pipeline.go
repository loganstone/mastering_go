package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			// fmt.Printf("\n duplicate! %d\n", x)
			CLOSEA = true
		} else {
			DATA[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x := range in {
		sum = sum + x
	}
	fmt.Printf("The sum of the random numbersi is %d\n", sum)
}

func main() {
	n1, n2 := 1, 100

	rand.Seed(time.Now().UnixNano())

	firstOutSecondIn := make(chan int)
	secondOut := make(chan int)

	go first(n1, n2, firstOutSecondIn)
	go second(secondOut, firstOutSecondIn)
	third(secondOut)
}
