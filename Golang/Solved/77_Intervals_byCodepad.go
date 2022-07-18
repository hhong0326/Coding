package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(intervals [][]int) {

	if len(intervals) <= 1 {
		fmt.Println(intervals)
		return
	}

	// sort start 기준
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	cur := intervals[0]
	res := [][]int{}

	for i := 1; i < len(intervals); i++ {
		s, e := cur[0], cur[1]
		is, ie := intervals[i][0], intervals[i][1]

		if is <= e {
			// update cur
			cur = []int{s, int(math.Max(float64(e), float64(ie)))}
		} else {
			res = append(res, cur)
			cur = intervals[i]
		}
	}

	// last element
	res = append(res, cur)

	fmt.Println(res)
}

func main() {

	solution([][]int{{1, 3}, {5, 8}, {4, 10}, {20, 25}})
}
