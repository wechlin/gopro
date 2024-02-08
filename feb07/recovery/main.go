package main

import (
	"fmt"
)

func get(values []float64, idx int) (res float64) {
	// Default value to return
	res = 1.11

	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
			fmt.Println("But wait, we'll recover from it!")
		}
	}()

	return values[idx]
}

func main() {
	data := []float64{3.14, 2.71, 6.02}

	fmt.Println(get(data, 2))
	fmt.Println(get(data, 0))
	fmt.Println(get(data, -1))
}
