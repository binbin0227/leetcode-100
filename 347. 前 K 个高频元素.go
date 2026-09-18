package main

import "container/heap"

type NumCount struct {
	num   int
	count int
}

type minHeap []NumCount

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}
func (h minHeap) Less(x, y int) bool {
	return h[x].count < h[y].count
}
func (h *minHeap) Push(x any) {
	*h = append(*h, x.(NumCount))
}
func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	for _, num := range nums {
		m[num]++
	}

	h := &minHeap{}
	heap.Init(h)

	for key, value := range m {
		heap.Push(h, NumCount{
			num:   key,
			count: value,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, 0, k)
	for _, item := range *h {
		res = append(res, item.num)
	}
	return res
}
