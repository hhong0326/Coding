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

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	var v int
	var arr []int

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v)
		arr = append(arr, v)
	}
	fmt.Fscanf(reader, "\n")

	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = arr[i] + sum[i]
	}

	var i, j int
	for k := 0; k < m; k++ {
		fmt.Fscanf(reader, "%d %d\n", &i, &j)

		answer := sum[j] - sum[i-1]

		fmt.Fprintln(writer, answer)
	}

}
