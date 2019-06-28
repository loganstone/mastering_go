// $ go test -benchmem -bench=. benchmark.go benchmark_test.go
package main

import "testing"

var result int

func benchmarkRecursion(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacciRecursion(n)
	}
	result = r
}

func benchmarkRecursionNonElse(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacciRecursionNonElse(n)
	}
	result = r
}

func benchmarkLoop(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = fibonacciLoop(n)
	}
	result = r
}

func Benchmark30FibonacciRecursion(b *testing.B) {
	benchmarkRecursion(b, 30)
}

func Benchmark30FibonacciRecursionNonElse(b *testing.B) {
	benchmarkRecursionNonElse(b, 30)
}

func Benchmark30FibonacciLoop(b *testing.B) {
	benchmarkLoop(b, 30)
}

func Benchmark50FibonacciRecursion(b *testing.B) {
	benchmarkRecursion(b, 50)
}

func Benchmark50FibonacciRecursionNonElse(b *testing.B) {
	benchmarkRecursionNonElse(b, 50)
}

func Benchmark50FibonacciLoop(b *testing.B) {
	benchmarkLoop(b, 50)
}
