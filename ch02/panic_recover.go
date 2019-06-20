package main

import (
	"fmt"
)

func recovering() {
	fmt.Println("Inside recovering function")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover inside recovering()")
		}
	}()

	fmt.Println("About to call panicking function")
	panicking()
	fmt.Println("panicking() exited")
	fmt.Println("Exiting recovering()")
}

func panicking() {
	fmt.Println("Inside panicking function")
	// log.Panic send message to stdout
	// panic not send message to stdout
	panic("Panic in panicking()")
	fmt.Println("Exiting panicking()")
}

func main() {
	recovering()
	fmt.Println("main() ended")
}
