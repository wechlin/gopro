package main

import (
	"fmt"
	"sync"
	"time"
)

type numeric interface {
	int | uint | float32 | float64 |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64
}

func generator[T any](done chan bool, vals ...T) <-chan T {
	src := make(chan T)

	go func() {
		defer close(src)

		for _, item := range vals {
			select {
			case src <- item:
			case <-done:
				return
			}
		}
	}()

	return src
}

func square[T numeric](done chan bool, in <-chan T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			select {
			case out <- v * v:
			case <-done:
				return
			}
		}
	}()

	return out
}

func merge[T any](done chan bool, ins ...<-chan T) <-chan T {
	var wg sync.WaitGroup

	out := make(chan T)

	slurp := func(in <-chan T) {
		defer wg.Done()

		for v := range in {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
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
	var closed = make(chan bool)
	source1 := generator(closed, 2, 4, 6)
	source2 := generator(closed, 3, 5, 7)

	stream := merge(closed, source1, source2)

	sq := square(closed, stream)

	fmt.Println(<-sq)
	fmt.Println(<-sq)
	fmt.Println(<-sq)
	fmt.Println(<-sq)

	close(closed)
	time.Sleep(1 * time.Second)
}
