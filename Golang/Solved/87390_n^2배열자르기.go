package main

import (
	"math"
)

func main() {
	solution(3, 2, 5)

}

func solution(n int, left int64, right int64) []int {

	// 1 222 33333 4444444 555555555
	// 1234 2234 3334 4444
	// 1 2 3
	// 2 2 3
	// 3 3 3
	// 1 2 3 2 2 3 3 3 3

	answer := []int{}

	for i := left; i <= right; i++ {
		d := float64(int(i) / n)
		m := float64(int(i) % n)

		answer = append(answer, int(math.Max(d, m)+1))
	}

	// for i := 1; i <= n; i++ {
	// 	// 1 2 3 4
	// 	for j := 0; j < i; j++ {

	// 		answer = append(answer, i)
	// 	}

	// 	for j := i; j < n; j++ {

	// 		answer = append(answer, j+1)

	// 	}
	// }

	return answer
	// return answer[left : right+1]
}
