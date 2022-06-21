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

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(a interface{}) { *h = append(*h, a.(int)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	v := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return v
}

func furthestBuilding(heights []int, bricks int, ladders int) int {

	// maxHeap

	//     h := &maxHeap{}

	//     LOOP:
	//     for i:=0; i<len(heights)-1; i++ {
	//         if heights[i+1] > heights[i] {
	//             diff := heights[i+1] - heights[i]
	//             heap.Push(h, diff)
	//             bricks -= diff
	//             if bricks >= 0 {
	//                 continue
	//             }

	//             for h.Len() > 0 {
	//                 maxDiff := heap.Pop(h).(int)
	//                 bricks += maxDiff
	//                 if bricks >= 0 {
	//                     ladders--
	//                     if ladders < 0 {
	//                         return i
	//                     } else {
	//                         continue LOOP
	//                     }
	//                 }
	//             }
	//             return i
	//         }
	//     }

	// minHeap

	h := &minHeap{}

	for i := 0; i < len(heights)-1; i++ {
		delta := heights[i+1] - heights[i]
		if delta <= 0 {
			continue
		}
		heap.Push(h, delta)
		if len(*h) > ladders {
			bricks -= heap.Pop(h).(int)
		}
		if bricks < 0 {
			return i
		}
	}

	return len(heights) - 1
}
