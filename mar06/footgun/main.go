package main

// Normal Go comment are anything in the file
// Only the comment block preceding the `import "C"` statement is compiled in C
// Any #cgo flag adjustments should be the first lines of the C comment block

// #cgo CPPFLAGS: -DSYMBOL
//     /*C Comment*/
//#include "header.h"
//#include <stdio.h>
//
//char message[] = "Hello World";
//
//int square(int x)
//{
//	return x*x;
//}
//
//int value = 7;
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	// Access C items in the C. pseudo-package
	var govalue = int(C.value)

	// C.int to conver Go ints to C ints
	fmt.Println(C.square(C.int(govalue)))

	fmt.Println(C.cube(C.int(govalue)))

	// Strings use CString() and GoString() functions for conversion
	C.puts(C.CString("Goodbye World"))

	// Arrays need their own conversion functions
	mess := C.GoBytes(unsafe.Pointer(&C.message), 12)
	fmt.Println(string(mess))
}
