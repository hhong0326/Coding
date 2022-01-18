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
	dp := []int{1, 2}

	for i := 3; i <= n; i++ {
		dp = append(dp, (dp[i-2]+dp[i-3])%15746)
	}

	fmt.Fprintln(writer, dp[n-1])
}
