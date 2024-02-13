package main

import "fmt"

type (
	Km = float64
	Kg float64
)

func arrayChunk() {
	var distance Km
	var length float64

	var legs = [...]Km{1100.0, 32.7, 55.5}

	copyLegs := legs

	distance = 1100.0
	length = 981.3

	for n, val := range copyLegs {
		fmt.Println(n, val)
	}

	legs[1]++
	fmt.Println(legs)
	fmt.Println(copyLegs)

	fmt.Println("Length is", len(copyLegs))

	if length == distance {
		fmt.Println("length and distance are equal")
	}

	fmt.Println(distance)
}

func mapChunk() {
	var grades map[string]float64

	if grades == nil {
		fmt.Println("Maps start out as nil")
	}

	grades = make(map[string]float64)

	key := "Bart"
	grades[key] = 70.0

	fmt.Println("Average for", key, "is", grades[key])
	key = "Bort"
	// Asking for a nonexistent key returns a 0-value
	fmt.Println("Average for", key, "is", grades[key])

	key = "Burt"
	val, ok := grades[key]
	if ok {
		fmt.Println("Found for", key, "is", val)
	} else {
		fmt.Println("Key", key, "not in grades")
	}

	delete(grades, "Bart")

	fmt.Println(grades)

	// Map literals separate key from value with colons
	grades = map[string]float64{"Bart": 70.0, "Milhouse": 81.3, "Nelson": 71.2}

	for k := range grades {
		fmt.Println(k)
	}
	for k, val := range grades {
		fmt.Println(k, ":", val)
	}
}

func main() {
	//arrayChunk()
	mapChunk()
}
