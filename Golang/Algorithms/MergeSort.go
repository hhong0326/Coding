package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

// recursive
func split(data []int) []int {

	fmt.Println("updated: ", data)
	left, right := []int{}, []int{}
	if len(data) <= 1 {
		return data
	} else {

		mid := len(data) / 2
		left, right = split(data[:mid]), split(data[mid:])
	}
	return merge(left, right)
}

func merge(left, right []int) []int {

	l, r := 0, 0
	// result := make([]int, 0, len(left)+len(right)+1)
	result := make([]int, 0)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	for l < len(left) {
		result = append(result, left[l])
		l++
	}

	for r < len(right) {
		result = append(result, right[r])
		r++
	}

	return result
}

func Start() {

	l := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		l[i] = rand.Intn(100-1) + 1
	}

	fmt.Println("before: ", l)

	result := split(l)

	fmt.Println("after: ", result)
}
