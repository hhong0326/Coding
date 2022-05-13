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

	// ways of stack
	type data struct {
		idx     int
		letters []int
	}

	s := []data{}
	output := [][]int{}

	s = append(s, data{0, []int{}})

	for len(s) > 0 {

		idx, letters := s[len(s)-1].idx, s[len(s)-1].letters
		s = s[:len(s)-1]

		if idx == len(nums) {
			output = append(output, letters)
			continue
		}

		c := nums[idx]

		s = append(s, data{idx + 1, letters})
		s = append(s, data{idx + 1, append(letters, c)})

	}
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

func main() {
	res := subsets([]int{1, 2, 3})
	fmt.Println(res)
}

// 참고: https://leetcode.com/problems/subsets/discuss/549771/go-backtrack
