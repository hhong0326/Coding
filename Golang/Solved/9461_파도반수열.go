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

	//  1, 1, 1, 2, 2, 3(1+2), 4(1+3), 5(1+4), 7(2+5), 9(2+7), 12(3+9), 16(4+12),

	var t, n int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		dp := make([]int, n+1)

		switch {
		case n <= 3:
			fmt.Fprintln(writer, 1)
			continue
		case 3 < n && n <= 5:
			fmt.Fprintln(writer, 2)
			continue
		}

		dp[1] = 1
		dp[2] = 1
		dp[3] = 1
		dp[4] = 2
		dp[5] = 2

		for i := 6; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-5]
		}

		fmt.Fprintln(writer, dp[n])
	}
}
