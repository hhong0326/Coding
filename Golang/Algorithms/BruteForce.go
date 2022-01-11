package main

import (
	"bufio"
	"fmt"
	"os"
)

// Brute Force
func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var w, h int
	var arr [][]int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &w, &h)
		arr = append(arr, []int{w, h})
	}

	var answer []int

	for i := 0; i < n; i++ {
		r := 1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if arr[i][0] < arr[j][0] && arr[i][1] < arr[j][1] {
				r++
			}
		}
		answer = append(answer, r)

	}

	for _, v := range answer {
		fmt.Fprintf(writer, "%d ", v)
	}

}
