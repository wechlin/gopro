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

func sliceChunk() {
	var fibs = [...]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	var slice []int = fibs[2:6]

	fmt.Println(slice)

	slice[3] = 33

	fmt.Println(slice)
	fmt.Println(fibs)

	// 90-95% of the time, this is how a slice is created and used
	// (via the make() builtin function)
	cats := make([]string, 0)

	// Making the underlying array have capacity means that initial appends
	// to cats will modify that initial underlying array, so if the slice
	// view that is returned is not replacing the original slice view,
	// unexpected things will happen!
	//cats := make([]string, 0, 4)

	fmt.Println("Length is", len(cats))

	fmt.Println("Cat capacity is", cap(cats))
	dogs := append(cats, "puma", "lynx")

	fmt.Println("cats:", cats)
	fmt.Println("dogs:", dogs)

	cats = append(cats, "leopardis", "uncia")

	fmt.Println("cats:", cats)
	fmt.Println("dogs:", dogs)

	slice = append(slice, 108, 99)

	fmt.Println(slice)
	fmt.Println(fibs)
}

func channelChunk() {
	var hopper chan string

	if hopper == nil {
		fmt.Println("Zero value for a channel is nil")
	}

	hopper = make(chan string, 3)

	// Sending is a statement by itself
	hopper <- "Complexity"

	fmt.Println(<-hopper + " is not so bad")

	sender(hopper)
	receiver(hopper)

	close(hopper)

	val, ok := <-hopper

	if !ok {
		fmt.Println("Channel is closed")
	} else {
		fmt.Println("val is", val)
	}
}

// Send-only channels
func sender(input chan<- string) {
	input <- "Value 1"
}

// Receive-only channels
func receiver(output <-chan string) {
	fmt.Println("In receiver:", <-output)

	// Cannot send to a receive-only channel
	//output <- "nope"
}

func main() {
	//arrayChunk()
	//mapChunk()
	//sliceChunk()
	channelChunk()
}
