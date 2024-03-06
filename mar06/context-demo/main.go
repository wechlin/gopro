package main

import (
	"fmt"
	"sync"
)

type numeric interface {
	int | uint | float32 | float64 |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64
}

func generator[T any](vals ...T) <-chan T {
	src := make(chan T)

	go func() {
		defer close(src)

		for _, item := range vals {
			src <- item
		}
	}()

	return src
}

func square[T numeric](in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()

	return out
}

func merge[T any](ins ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	out := make(chan T)

	slurp := func(in <-chan T) {
		for v := range in {
			out <- v
		}
		wg.Done()
	}

	wg.Add(len(ins))
	go func() {

		for _, in := range ins {
			go slurp(in)
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	source1 := generator(2, 4, 6)
	source2 := generator(3, 5, 7)

	stream := merge(source1, source2)

	sq := square(stream)

	fmt.Println(<-sq)
	fmt.Println(<-sq)
	fmt.Println(<-sq)
	fmt.Println(<-sq)
}
