package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
// //

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)

// 	defer writer.Flush()

// 	var n int

// 	fmt.Fscanln(reader, &n)
// 	var sue int
// 	var arrMin IntHeap
// 	var arrMax IntHeap

// 	for n > 0 {

// 		fmt.Fscanln(reader, &sue)

// 		if len(arrMin) == 0 || arrMin[len(arrMin)-1] >= sue {

// 			arrMin = append(arrMin, sue)
// 			if len(arrMin) > len(arrMax)+1 {
// 				arrMax.Push(sue)

// 				arrMin.Pop()

// 			}
// 		} else if len(arrMax) == 0 || arrMax[len(arrMax)-1] <= sue {
// 			arrMax = append(arrMax, sue)
// 			if len(arrMax) > len(arrMin) {
// 				arrMin.Push(sue)
// 				arrMax.Pop()
// 			}
// 		} else {
// 			if len(arrMin) <= len(arrMax) {
// 				arrMin.Push(sue)
// 			} else {
// 				arrMax.Push(sue)
// 			}
// 		}

// 		fmt.Println(arrMin)
// 		fmt.Println(arrMax)

// 		n--

// 		fmt.Fprintln(writer, arrMin[len(arrMin)-1])

// 	}
// }

// 시간 초과
func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int

	fmt.Fscanln(reader, &n)
	var sue int
	var arr []int
	for n > 1 {

		fmt.Fscanln(reader, &sue)
		arr = append(arr, sue)

		if len(arr) <= 1 {
			fmt.Fprintln(writer, arr[0])
			continue
		}

		sort.Ints(arr)

		if len(arr)%2 == 0 {
			fmt.Fprintln(writer, arr[len(arr)/2-1])

		} else {
			fmt.Fprintln(writer, arr[len(arr)/2])
		}

		n--
	}
}
