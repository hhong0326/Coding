package main

import (
	"fmt"
	"sort"
	"strconv"
)

func combinationSum2(candidates []int, target int) [][]int {

	res := [][]int{}

	sort.Ints(candidates)
	m := make(map[string]bool, 0)

	backtrack(0, candidates, &res, target, "", []int{}, m)

	return res
}

func backtrack(idx int, candidates []int, res *[][]int, t int, s string, arr []int, m map[string]bool) {

	if _, ok := m[s]; ok {
		return
	}
	m[s] = true
	if sumArr(arr) == t {
		fmt.Println(arr)
		*res = append(*res, arr)
		return
	} else if sumArr(arr) > t {
		return
	}

	for i := idx; i < len(candidates); i++ {

		backtrack(i+1, candidates, res, t, s+strconv.Itoa(candidates[i]), append(arr, candidates[i]), m)

	}

}

func sumArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
