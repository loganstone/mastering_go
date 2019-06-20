package main

// #include <stdio.h>
// void callC() {
//     printf("Calling C code!\n");
// }
import "C" // Must be no new line after c code.
import "fmt"

func main() {
	fmt.Println("A Go statement")
	C.callC()
	fmt.Println("Another Go statement")
}
