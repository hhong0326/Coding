package main

func maximumUniqueSubarray(nums []int) int {

	s, sum, max := 0, 0, 0
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		if idx, ok := m[nums[i]]; ok {
			for j := s; j <= idx; j++ {
				sum -= nums[j]
				delete(m, nums[j])
			}
			s = idx + 1
		}

		m[nums[i]] = i
		sum += nums[i]
		if max < sum {
			max = sum
		}

	}

	return max

	// other's solution
	// l, sum, max := 0, 0, 0
	// m := map[int]int{}
	// for r := 0; r < len(nums); r++ {
	// 	if j, ok := m[nums[r]]; ok && j >= l {
	// 		for ; l <= j; l++ {
	// 			sum -= nums[l]
	// 		}
	// 	}
	// 	sum += nums[r]
	// 	m[nums[r]] = r
	// 	if sum > max {
	// 		max = sum
	// 	}
	// }
	// return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
