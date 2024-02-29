package priorityqueue

import "testing"

func TestSize(t *testing.T) {
	empty := MinHeap[int, int]{}

	if sz := empty.Size(); sz != 0 {
		t.Errorf(".Size() = %d, want 0", sz)
	}

	if _, _, err := empty.Dequeue(); err == nil {
		t.Errorf("{}.Dequeue() did not error")
	}

}

func TestEnqueuing(t *testing.T) {
	pq := MinHeap[int, int]{}

	pq.Enqueue(5, 2)
	pq.Enqueue(0, 1)

	p, o, err := pq.Dequeue()
	if err != nil {
		t.Errorf("Dequeue() errored, %v", err)
	}
	if p != 0 || o != 1 {
		t.Errorf("Dequeue() = %d, %d, want %d, %d", p, o, 0, 1)
	}

	p, o, err = pq.Dequeue()
	if err != nil {
		t.Errorf("Dequeue() errored, %v", err)
	}
	if p != 5 || o != 2 {
		t.Errorf("Dequeue() = %d, %d, want %d, %d", p, o, 5, 2)
	}
}
