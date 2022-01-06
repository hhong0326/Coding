package main

import (
	"bufio"
	"fmt"
	"os"
)

// DP

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var testCaseLength int
	fmt.Fscanln(reader, &testCaseLength)

	var n int
	for i := 0; i < testCaseLength; i++ {
		fmt.Fscanln(reader, &n)

		dp := make([]int, n+1)

		dp[0] = 0

		if n == 1 {
			fmt.Fprintln(writer, n)
			writer.Flush()
		} else if n == 2 {
			fmt.Fprintln(writer, n)
			writer.Flush()
		} else if n == 3 {
			fmt.Fprintln(writer, n+1)
			writer.Flush()
		} else {

			dp[1] = 1
			dp[2] = 2
			dp[3] = 4

			for i := 4; i <= n; i++ {
				dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
			}

			fmt.Fprintln(writer, dp[n])
			writer.Flush()
		}

	}
}
