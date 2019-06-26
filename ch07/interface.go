package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

// Read the argument type.
func calculate(x shape) {
	switch v := x.(type) {
	case square:
		fmt.Println("Is a square", v)
		fmt.Println(x.Area())
		fmt.Println(x.Perimeter())
	case circle:
		fmt.Println("Is a circle!")
		fmt.Println(x.Area())
		fmt.Println(x.Perimeter())
	default:
		fmt.Printf("Etc %T\n", v)
	}
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("This is a square", v)
	case circle:
		fmt.Printf("%v is a circle!\n", v)
	case rectangle:
		fmt.Println("This is a rectangle")
	default:
		fmt.Printf("Unknow type %T!\n", v)
	}
}

func main() {
	x := square{X: 10}
	fmt.Println("Perimeter:", x.Perimeter())
	calculate(x)

	y := circle{R: 5}
	calculate(y)

	tellInterface(x)
	tellInterface(y)
	tellInterface(rectangle{})
	tellInterface(10)
}
