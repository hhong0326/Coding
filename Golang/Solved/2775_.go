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

	var n int
	fmt.Fscanln(r, &n)

	for i := 0; i < n; i++ {

		var k, m int

		fmt.Fscanln(r, &k)
		fmt.Fscanln(r, &m)

		dp := make([][]int, k+1)
		for i := 0; i < len(dp); i++ {
			dp[i] = make([]int, m)
		}

		for i := 0; i < m; i++ {
			dp[0][i] = i + 1
		}

		for i := 1; i < len(dp); i++ {
			var sum int
			for j := 0; j < m; j++ {
				for k := 0; k <= j; k++ {

					sum += dp[i-1][k]
				}
				dp[i][j] = sum
				sum = 0

			}
		}
		fmt.Fprintln(w, dp[k][m-1])

	}
}
