package main

import "container/heap"

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MaxHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MedianFinder struct {
	left  *MaxHeap
	right *MinHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &MaxHeap{},
		right: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.left, num)
	heap.Push(this.right, heap.Pop(this.left))

	// 保证左右一样多，或者左边比右边多一个。由于每一次添加都遵守这个规则，所以不需要用for
	if this.left.Len() < this.right.Len() {
		heap.Push(this.left, heap.Pop(this.right))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if ((*this.left).Len()+(*this.right).Len())%2 == 0 {
		return (float64((*this.left)[0]) + float64((*this.right)[0])) / 2
	} else {
		return float64((*this.left)[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
