package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {

	// m := make(map[string]bool, 0)

	// res := [][]int{}

	// fmt.Println(res)

	// sort.Ints(nums)

	// backtrack(res, nums, []int{}, 0, m)
	// return res

	res := make([][]int, 0)

	backtrack(0, nums, nil, &res)
	return res
}

func backtrack(index int, nums []int, cur []int, res *[][]int) {
	if index == len(nums) {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	backtrack(index+1, nums, cur, res)
	backtrack(index+1, nums, append(cur, nums[index]), res)
}

// func backtrack(res [][]int, nums []int, arr []int, cnt int, m map[string]bool) {
// 	if cnt == len(nums) {
// 		res = append(res, arr)
// 		return
// 	}

// 	var str string
// 	for _, v := range arr {
// 		s := strconv.Itoa(v)
// 		str += s
// 	}

// 	if _, ok := m[str]; !ok {
// 		res = append(res, arr)
// 		m[str] = true
// 	}

// 	arr = append(arr, nums[cnt])
// 	fmt.Println(arr)
// 	cnt += 1

// 	backtrack(res, nums, arr, cnt, m)

// }

func main() {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)
}

// 참고: https://leetcode.com/problems/subsets/discuss/549771/go-backtrack
