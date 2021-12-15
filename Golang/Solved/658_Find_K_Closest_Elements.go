package coding

// Binary search way

// left < right
// nums[mid] < target
// left = mid + 1

func findClosestElements(arr []int, k int, x int) []int {

	s, e := 0, len(arr)-k

	for s+1 < e {
		mid := (s + e) / 2
		if x-arr[mid] > arr[mid+k]-x {
			s = mid
		} else {
			e = mid
		}
	}
	return arr[s : s+k]
}
