package main

import (
	"fmt"
	"unsafe"
)

type thing struct {
	Name string
}

func main() {
	var ref uintptr = 5
	var ptr = new(uint)
	*ptr = 7

	var obj = new(thing)
	obj.Name = "Andy"

	ref += 5

	var up = unsafe.Pointer(ptr)
	up = unsafe.Pointer(obj)
	ref = uintptr(up)
	ref += 8
	// The linter ain't lyin', this is dangerous
	up = unsafe.Pointer(ref)

	var val = (*int)(up)

	fmt.Println(ref, ptr, up, obj, *val)
}
