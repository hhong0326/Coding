package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	var max int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if i == j || j == k || k == i {
					continue
				}

				v := arr[i] + arr[j] + arr[k]
				if m >= v {
					max = int(math.Max(float64(max), float64(v)))
				}
			}
		}
	}

	fmt.Fprintln(writer, max)
}
