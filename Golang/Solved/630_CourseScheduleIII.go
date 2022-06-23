package main

import (
	"container/heap"
	"sort"
)

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

func scheduleCourse(courses [][]int) int {

	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := &maxHeap{}
	t := 0
	for _, course := range courses {
		heap.Push(h, course[0])
		t += course[0]

		if t > course[1] {
			t -= heap.Pop(h).(int)
		}

		if t > course[1] {
			return h.Len()
		}
	}

	return h.Len()
}
