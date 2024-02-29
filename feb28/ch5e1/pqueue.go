package priorityqueue

import (
	"cmp"
	"errors"
)

type PriorityQueue[P cmp.Ordered, T any] interface {
	Enqueue(pri P, obj T)
	Dequeue() (P, T, error)
	Size() int
}

type pair[P cmp.Ordered, T any] struct {
	priority P
	object   T
}
type MinHeap[P cmp.Ordered, T any] []pair[P, T]

func (mh *MinHeap[P, T]) Size() int {
	return len(*mh)
}

func (mh *MinHeap[P, T]) Enqueue(pri P, obj T) {
	*mh = append(*mh, pair[P, T]{pri, obj})
	(*mh).bubbleUp(len(*mh) - 1)
}

func (mh *MinHeap[P, T]) Dequeue() (pri P, obj T, err error) {
	if len(*mh) == 0 {
		err = errors.New("empty Heap")
		return
	}
	pri, obj = (*mh)[0].priority, (*mh)[0].object

	(*mh).bubbleDown()
	*mh = (*mh)[:len(*mh)-1]
	return
}

func (mh MinHeap[P, T]) bubbleDown() {
	idx := 0

	mh[0] = mh[len(mh)-1]

	for idx < len(mh)-1 {
		left := 2*idx + 1
		right := 2*idx + 2
		toSwap := idx
		if left < len(mh) && mh[left].priority < mh[toSwap].priority {
			toSwap = left
		}
		if right < len(mh) && mh[right].priority < mh[toSwap].priority {
			toSwap = right
		}

		if idx != toSwap {
			mh[idx], mh[toSwap] = mh[toSwap], mh[idx]
		} else {
			return
		}
		idx = toSwap
	}

}

func (mh MinHeap[P, T]) bubbleUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if mh[idx].priority < mh[parentIdx].priority {
			mh[idx], mh[parentIdx] = mh[parentIdx], mh[idx]

		} else {
			return
		}
		idx = parentIdx
	}
}
