package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c)

	go func() {
		for {
			sig := <-c
			switch sig {
			case os.Interrupt:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall.SIGUSR2")
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep(30 * time.Second)
	}
}
