package main

import "fmt"

// Dog .
type Dog struct {
	Name string
	Age  int
}

func main() {
	// return value
	dogs := make([]Dog, 0)
	fmt.Println(dogs)
	fmt.Println(&dogs)

	// return pointer
	pointerDogs := new([]Dog)
	fmt.Println(pointerDogs)
	fmt.Println(*pointerDogs)

	pointerDog := new(Dog)
	fmt.Println(pointerDog)
	fmt.Println(*pointerDog)
}
