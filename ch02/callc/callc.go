// Excute
// $ gcc -c lib/*.c
// $ ar rs callc.a *.o
// $ go run callc.go

package main

// #cgo CFLAGS: -I${SRCDIR}/lib
// #cgo LDFLAGS: ${SRCDIR}/callc.a
// #include <stdlib.h>
// #include <callc.h>
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("A Go statement")
	C.cHello()

	fmt.Println("Another Go statement")
	message := C.CString("this is Mihalis")
	defer C.free(unsafe.Pointer(message))
	C.printMessage(message)

	fmt.Println("Done")
}
