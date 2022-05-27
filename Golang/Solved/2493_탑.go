package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var N int
	fmt.Fscanln(r, &N)

	var arr []int
	for i := 0; i < N; i++ {
		var m int
		fmt.Fscanf(r, "%d ", &m)
		arr = append(arr, m)
	}

	// fmt.Println(arr)

	// 탑의 번호 = 인덱스 + 1

	res := make([]int, len(arr))

	// for i := len(arr) - 1; i >= 1; i-- {

	// 	for j := i - 1; j >= 0; j-- {

	// 		if arr[i] < arr[j] {
	// 			res[i] = j + 1
	// 			break
	// 		}
	// 	}
	// }

	// sliding window

	// s, e := len(arr)-1, len(arr)-2

	// for s > 0 {

	// 	// 수신
	// 	if arr[s] < arr[e] {
	// 		res[s] = e + 1
	// 		s-- // 다음 탑으로

	// 		e = s - 1 // 다음 탑의 다음부터
	// 	} else {
	// 		// 다음 탑으로
	// 		e--
	// 	}
	// }

	// front first

	for i := 1; i < N; i++ {
		prev := i - 1
		for prev >= 0 {
			if arr[i] == arr[prev] {
				res[i] = res[prev] // 수신한 탑의 번호를 송신 탑에 씀
				break
			} else if arr[i] > arr[prev] { // 이전 수신한 탑과 다시 비교한다
				prev = res[prev] - 1
			} else { // arr[i] < arr[prev] // 현재 탑은 이전 탑이 수신한다
				res[i] = prev + 1
				break
			}

		}
	}

	for _, v := range res {

		fmt.Fprintf(w, "%d ", v)
	}
}

// Reference
// https://deok2kim.tistory.com/246
