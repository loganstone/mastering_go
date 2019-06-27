package main

import (
	"flag"
	"fmt"
	"sync"
)

var n = flag.Int("n", 20, "Number of goroutines")

func main() {
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	wg.Wait()
	fmt.Println("\nExiting...")
}
