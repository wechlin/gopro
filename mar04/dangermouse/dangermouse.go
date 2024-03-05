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

	var val = (*int)(unsafe.Add(unsafe.Pointer(obj), unsafe.Sizeof(obj.Name)))

	fmt.Println(ref, ptr, *val)
}
