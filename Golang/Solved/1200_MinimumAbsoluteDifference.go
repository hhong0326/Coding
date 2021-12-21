package solved

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {

	result := make([][]int, 0)

	sort.Ints(arr)

	min := 1000000

	for i := 0; i < len(arr)-1; i++ {
		abs := math.Abs(float64(arr[i] - arr[i+1]))
		if float64(min) > abs {
			min = arr[i+1] - arr[i]
			result = make([][]int, 0)
			result = append(result, []int{arr[i], arr[i+1]})
		} else if float64(min) == abs {
			result = append(result, []int{arr[i], arr[i+1]})
		}
	}

	return result
}
