package coding

import "fmt"

// 128
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	for _, num := range nums {
		set[num] = struct{}{}
	}

	max, count := 0, 0

	for num := range set {
		if _, ok := set[num-1]; ok {
			fmt.Println(num)
			continue
		}

		current := num
		for {
			if _, ok := set[num]; ok {
				current++
				count++
			} else {
				if max < count {
					max = count
				}
				count = 0

				break
			}
		}
	}

	return max
}
