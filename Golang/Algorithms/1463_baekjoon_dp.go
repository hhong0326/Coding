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

	var testCaseLength int
	fmt.Fscanln(reader, &testCaseLength)

	var n int
	fmt.Fscanln(reader, &n)

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/2]+1)))
		}
		if i%3 == 0 {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i/3]+1)))
		}
	}

	fmt.Fprintln(writer, dp[n])

	writer.Flush()
}
