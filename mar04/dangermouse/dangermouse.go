package main

import (
	"fmt"
	"unsafe"
)

type thing struct {
	Name string
	age  int
}

func main() {
	var ref uintptr = 5
	var ptr = new(uint)
	*ptr = 7

	var obj = new(thing)
	obj.Name = "Andy"
	obj.age = 45

	ref += 5

	var up = unsafe.Pointer(ptr)
	up = unsafe.Pointer(obj)
	ref = uintptr(up)
	ref += unsafe.Sizeof(obj.Name)
	// The linter ain't lyin', this is dangerous
	up = unsafe.Pointer(ref)

	var val = (*int)(up)

	fmt.Println(ref, ptr, up, obj, *val)
}
