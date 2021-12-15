package coding

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, start, end, k int) int {
	if start == end {
		return nums[start]
	}

	pivot := (start + end) / 2
	pivot = partition2(nums, start, end, pivot)

	if k > pivot {
		return quickSelect(nums, pivot+1, end, k)
	} else if k < pivot {
		return quickSelect(nums, start, pivot-1, k)
	} else {
		return nums[pivot]
	}
}

func partition2(nums []int, start, end, pivot int) int {
	swap(nums, end, pivot)
	store := start
	for i := start; i <= end; i++ {
		if nums[i] < nums[end] {
			swap(nums, store, i)
			store++
		}
	}
	swap(nums, store, end)
	return store
}

func swap(nums []int, i, j int) {
	tmp := nums[j]
	nums[j] = nums[i]
	nums[i] = tmp
}
