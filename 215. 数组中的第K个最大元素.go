package main

import "container/heap"

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(x, y int) bool {
	return h[x] < h[y]
}

func (h minHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)

		if len(*h) > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
