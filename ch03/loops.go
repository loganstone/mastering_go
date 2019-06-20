package main

import (
	"fmt"
)

func main() {
	// normally
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// while
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()

	// do..while
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Print(i, " ")
		i++
	}
	fmt.Println()

	// range
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, v := range anArray {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}
