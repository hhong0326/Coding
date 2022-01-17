package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {

		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	var m int
	fmt.Fscanln(reader, &m)

	var s, e int

	questions := make([][]int, m)
	for i := 0; i < m; i++ {

		fmt.Fscanf(reader, "%d %d\n", &s, &e)
		questions[i] = append(questions[i], s, e)
	}

	var answer []int
	for i := 0; i < m; i++ {
		answer = append(answer, isPalin(arr[questions[i][0]-1:questions[i][1]]))
	}

	for _, v := range answer {
		fmt.Fprintln(writer, v)
	}
}

func isPalin(arr []int) int {

	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return 0
		}
	}
	return 1
}
