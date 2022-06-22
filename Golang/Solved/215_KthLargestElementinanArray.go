package main

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(a interface{}) { *h = append(*h, a.(int)) }
func (h *maxHeap) Pop() interface{} {
	l := len(*h)
	v := (*h)[l-1]
	*h = (*h)[:l-1]
	return v
}

func findKthLargest(nums []int, k int) int {

	h := &maxHeap{}

	for _, v := range nums {
		heap.Push(h, v)
	}

	for k > 1 {
		heap.Pop(h)
		k--
	}

	return heap.Pop(h).(int)

	// sort.Ints(nums)
	// return nums[len(nums)-k]
}
