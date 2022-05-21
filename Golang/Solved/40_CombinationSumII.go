package main

import (
	"sort"
)

// func combinationSum2(candidates []int, target int) [][]int {

// 	res := [][]int{}

// 	sort.Ints(candidates)
// 	m := make(map[string]bool, 0)

// 	backtrack(0, candidates, &res, target, "", []int{}, m)

// 	return res
// }

// func backtrack(idx int, candidates []int, res *[][]int, t int, s string, arr []int, m map[string]bool) {

// 	if _, ok := m[s]; ok {
// 		return
// 	}
// 	m[s] = true
// 	if sumArr(arr) == t {
// 		fmt.Println(arr)
// 		*res = append(*res, arr)
// 		return
// 	} else if sumArr(arr) > t {
// 		return
// 	}

// 	for i := idx; i < len(candidates); i++ {

// 		backtrack(i+1, candidates, res, t, s+strconv.Itoa(candidates[i]), append(arr, candidates[i]), m)

// 	}

// }

// func sumArr(arr []int) int {
// 	sum := 0
// 	for _, v := range arr {
// 		sum += v
// 	}

// 	return sum
// }

// Others solution
var res [][]int
var candis []int

func combinationSum2(candidates []int, target int) [][]int {
	res = nil
	candis = candidates

	sort.Ints(candidates)
	backtrack(target, nil, 0)
	return res
}

func backtrack(target int, cur []int, index int) {
	if target == 0 {
		res = append(res, append([]int{}, cur...))
		return
	}

	if target < 0 {
		return
	}

	for i := index; i < len(candis); i++ {
		if i > index && candis[i] == candis[i-1] {
			continue
		}
		backtrack(target-candis[i], append(cur, candis[i]), i+1)
	}
}
