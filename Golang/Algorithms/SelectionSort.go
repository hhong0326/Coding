package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

func selectionSort(nums []int) []int {

	// 마지막 원소는 자동정렬됨
	for i := 0; i < len(nums)-1; i++ {

		least := i
		// 최솟값 탐색
		for j := i + 1; j < len(nums); j++ {
			if nums[least] > nums[j] {
				least = j
			}
			// fmt.Println(nums[least])
		}

		// 자기 자신 이동 x
		if i != least {
			nums[i], nums[least] = nums[least], nums[i]
		}
	}

	return nums
}

func StartSelectSort() {

	l := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {

		l[i] = rand.Intn(100-1) + 1
	}

	fmt.Println("before: ", l)
	fmt.Println("after: ", selectionSort(l))
}
