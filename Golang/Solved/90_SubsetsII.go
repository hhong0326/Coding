package main

import (
	"sort"
	"strconv"
)

func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	m := map[string][]int{}
	backtrack([]int{}, nums, "", m)
	res := make([][]int, len(m))
	i := 0
	for _, s := range m {
		res[i] = s
		i++
	}
	return res
}

func backtrack(current, lefts []int, s string, m map[string][]int) {
	if _, ok := m[s]; ok {
		return
	}

	m[s] = current

	if len(lefts) == 0 {
		return
	}

	for i, r := range lefts {
		nextCurrent := make([]int, len(current))
		copy(nextCurrent, current)
		nextCurrent = append(nextCurrent, r)
		backtrack(nextCurrent, lefts[i+1:], s+strconv.Itoa(r), m)
	}
}
