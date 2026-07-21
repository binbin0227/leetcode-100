package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for i := range len(nums) {
		if h.Len() < k {
			// 前 k 个人直接进堆
			heap.Push(h, nums[i])
		} else {
			if nums[i] > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, nums[i])
			}
		}
	}
	return (*h)[0]
}
