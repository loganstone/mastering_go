package main

import "fmt"

func fibonacciRecursion(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacciRecursion(n-1) + fibonacciRecursion(n-2)
	}
}

func fibonacciRecursionNonElse(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacciRecursionNonElse(n-1) + fibonacciRecursionNonElse(n-2)
}

func fibonacciLoop(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}
		fn[i] = f
	}
	return fn[n]
}

func main() {
	fmt.Println(fibonacciRecursion(40))
	fmt.Println(fibonacciRecursionNonElse(40))
	fmt.Println(fibonacciLoop(40))
}
