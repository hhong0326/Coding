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

	var n, k int
	fmt.Fscanf(reader, "%d %d\n", &n, &k)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &arr[i])
	}

	var count int
	i := len(arr) - 1
	for k != 0 {

		if k >= arr[i] {
			v := k / arr[i]
			k %= arr[i]
			count += v
		} else {
			i--
		}
	}

	fmt.Fprintln(writer, count)
}
