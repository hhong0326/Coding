package main

func permuteUnique(nums []int) [][]int {

	res := [][]int{}
	backtrackSwap(&res, 0, nums)

	return res
}

func backtrackSwap(ans *[][]int, depth int, nums []int) {
	if depth == len(nums) {
		*ans = append(*ans, append([]int{}, nums...))
	}

	used := make(map[int]int)
	for i := depth; i < len(nums); i++ {
		if _, ok := used[nums[i]]; ok {
			continue
		}

		nums[i], nums[depth] = nums[depth], nums[i]
		backtrackSwap(ans, depth+1, nums)
		nums[i], nums[depth] = nums[depth], nums[i]

		used[nums[i]] = i
	}
}
