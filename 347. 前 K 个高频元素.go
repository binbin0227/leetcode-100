package main

import "container/heap"

type Card struct {
	num   int
	count int
}

type MinHeap []Card

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].count < h[j].count
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Card))
}
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	h := &MinHeap{}
	heap.Init(h)

	for num, count := range countMap {
		if h.Len() < k {
			heap.Push(h, Card{
				num:   num,
				count: count,
			})
		} else {
			if count > (*h)[0].count {
				heap.Pop(h)
				heap.Push(h, Card{
					num:   num,
					count: count,
				})
			}
		}
	}

	res := make([]int, 0, k)
	for _, card := range *h {
		res = append(res, card.num)
	}
	return res
}
