package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

func quickSortV2(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortV2(arr, low, p-1)
		arr = quickSortV2(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func StartQuick() {

	l := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		l[i] = rand.Intn(100-1) + 1
	}

	fmt.Println("before: ", l)

	start2 := time.Now()
	result := quickSortV2(l, 0, len(l)-1)
	fmt.Println(time.Since(start2))
	fmt.Println("after v2: ", result)
}
