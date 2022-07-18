package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InsertSort(arr []int) []int {

	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // 오른쪽 이동
			j--
		}
		arr[j+1] = key
	}
	return arr
}

func main() {
	arr := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		arr[i] = rand.Intn(100-1) + 1
	}
	fmt.Println(arr)
	fmt.Println(InsertSort(arr))
}

// https://gmlwjd9405.github.io/2018/05/06/algorithm-insertion-sort.html
