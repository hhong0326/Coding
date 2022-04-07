package main

func main() {
	solution(3, 2, 5)

}

func solution(n int, left int64, right int64) []int {

	// 1 222 33333 4444444 555555555
	// 1234 2234 3334 4444
	// 1 2 3
	// 2 2 3
	// 3 3 3

	// arr := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	arr[i] = make([]int, n)
	// }

	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		arr[i][j] = i + 1
	// 		arr[j][i] = i + 1
	// 	}
	// }

	answer := []int{}

	// for i := 0; i < n; i++ {
	// 	answer = append(answer, arr[i]...)
	// }

	// find pattern
	// 1 222 33333 4444444 555555555
	// 1234 2234 3334 4444

	for i := 1; i <= n; i++ {
		// 1 2 3 4
		for j := 0; j < i; j++ {

			answer = append(answer, i)
		}

		for j := i; j < n; j++ {

			answer = append(answer, j+1)

		}
	}

	// fmt.Println(answer[left : right+1])

	return answer[left : right+1]
}
