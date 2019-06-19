package main

import (
	"errors"
	"fmt"
	"log"
)

const errMsg = "error in newError function"

func newError(a, b int) error {
	if a == b {
		return errors.New(errMsg)
	}
	return nil
}

func main() {
	err := newError(1, 2)
	if err == nil {
		fmt.Println("ended normally")
	} else {
		fmt.Println(err)
	}

	err = newError(10, 10)
	if err == nil {
		fmt.Println("ended normally")
	} else {
		log.Println(err)
		// Exit immediate
		panic(err)
	}

	if err.Error() == errMsg {
		fmt.Println("!!")
	}
}
