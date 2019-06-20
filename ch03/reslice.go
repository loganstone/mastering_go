package main

import "fmt"

func main() {
	// slice is [0, 0, 0, 0, 0]
	slice := make([]int, 5)
	fmt.Println(slice)
	// refer, not copy
	reSlice := slice[1:3]
	fmt.Println(reSlice)
	// slice is also changed
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(slice)
	fmt.Println(reSlice)
}
