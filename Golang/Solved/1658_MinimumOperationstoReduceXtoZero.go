package main

func minOperations(nums []int, x int) int {
	target := -x

	for _, v := range nums {
		target += v
	}

	if target == 0 {
		return len(nums)
	}

	if target < 0 {
		return -1
	}

	res := -1

	s, e, sum := 0, 0, 0
	for e < len(nums) {
		sum += nums[e]
		e++
		for sum > target && s < e {
			sum -= nums[s]
			s++
		}
		if target == sum && e-s > res {
			res = e - s
		}
	}

	if res == -1 {
		return res
	}

	return len(nums) - res
}

// Reference
// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/discuss/1017353/Go-prefixSum-%2B-HashMap-VS-Sliding-window~
