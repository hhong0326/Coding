package main

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {

	sort.Ints(nums)

	med := nums[len(nums)/2]

	sum := 0
	for _, v := range nums {
		sum += int(math.Abs(float64(med - v)))
	}

	return sum
}
