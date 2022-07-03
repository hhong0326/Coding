package main

func wiggleMaxLength(nums []int) int {

	//     dpUp := make([]int, len(nums))
	//     dpDown := make([]int, len(nums))

	//     dpUp[0], dpDown[0] = 1, 1

	up, down := 1, 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// dpUp[i] = dpDown[i-1] + 1
			// dpDown[i] = dpDown[i-1]
			up = down + 1
		} else if nums[i] < nums[i-1] {
			// dpDown[i] = dpUp[i-1] + 1
			// dpUp[i] = dpUp[i-1]
			down = up + 1
		}
	}

	return max(up, down)

	// return max(dpUp[len(dpUp)-1], dpDown[len(dpDown)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
