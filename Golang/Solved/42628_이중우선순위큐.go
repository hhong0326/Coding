package main

import (
	"container/heap"
	"strconv"
	"strings"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func solution(operations []string) []int {

	h := &IntHeap{} // min heap
	// q := []int{}

	heap.Init(h)
	for _, op := range operations {
		command := strings.Split(op, " ")

		switch command[0] {
		case "I":
			v, _ := strconv.Atoi(command[1])

			// q = append(q, v)
			// sort.Ints(q)
			heap.Push(h, v)

		case "D":
			// if len(q) == 0 {
			// 	continue
			// }
			if h.Len() == 0 {
				continue
			}
			if command[1] == "1" {
				// delete max
				// q = q[:len(q)-1]
				heap.Remove(h, h.Len()-1)
			} else {
				// delete min
				// q = q[1:]
				heap.Pop(h) // min Heap
			}
		}
	}

	answer := []int{0, 0}
	// if len(q) == 0 {
	// 	return answer
	// }

	if h.Len() == 0 {
		return answer
	}

	checkIndex := 0
	//0 최대 // 1 최소값
	for h.Len() > 0 {
		val := heap.Pop(h)

		if checkIndex == 0 {
			//First 최소값
			answer[1] = val.(int)
			checkIndex = 1
		} else if h.Len() == 0 {
			//Last 최대값
			answer[0] = val.(int)
		}
	}

	// fmt.Println(q[len(q)-1], q[0])
	// fmt.Println((*h)[h.Len()-1], (*h)[0])
	// return []int{q[len(q)-1], q[0]}

	return answer
}

func main() {
	solution([]string{"I 16", "D 1"})
	solution([]string{"I 7", "I 5", "I -5", "D -1"})
}

// Reference
// https://programmers.co.kr/learn/courses/30/lessons/42628/solution_groups?language=go
