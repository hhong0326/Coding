package main

func findMin(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	// O(logN)
	s, e := 0, len(nums)-1

	if nums[e] > nums[0] {
		return nums[0]
	}

	for s < e {
		mid := (s + e) / 2

		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}

		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}

		if nums[mid] > nums[0] {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return nums[0]

	// O(n)
	// s := 0
	// for i, v := range nums {
	//     if nums[s] > v {
	//         return nums[i]
	//     }
	// }
	// return nums[s]

}
