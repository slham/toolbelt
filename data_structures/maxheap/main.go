package main

import "log"

// MaxHeap
type MaxHeap struct {
	array []int
}

// Insert
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract
func (h *MaxHeap) Extract() int {
	out := h.array[0]
	l := len(h.array) - 1

	if l == -1 {
		log.Println("empty array")
		return l
	}

	h.array[0] = h.array[l]
	h.array = h.array[:l]

	return out
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= lastIndex {
		if l == lastIndex { // left child is only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // left child is larger
			childToCompare = l
		} else { // right child is larger
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	maxHeap := &MaxHeap{}
	log.Println(maxHeap)

	for _, val := range []int{1, 1, 2, 3, 5, 8, 13, 21} {
		maxHeap.Insert(val)
		log.Println(maxHeap)
	}

	for i := 0; i <= 5; i++ {
		log.Println("extracting", maxHeap.Extract())
		log.Println(maxHeap)
	}
}
