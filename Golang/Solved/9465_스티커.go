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

	var n int
	fmt.Fscanln(reader, &n)

	for k := 0; k < n; k++ {
		var m int
		fmt.Fscanln(reader, &m)
		sticker := make([][]int, 2)
		dp := make([][]int, 2)
		for i := 0; i < 2; i++ {
			sticker[i] = make([]int, m)
			dp[i] = make([]int, m)
			for j := 0; j < m; j++ {
				fmt.Fscanf(reader, "%d ", &sticker[i][j])
			}
		}
		dp[0][0] = sticker[0][0]
		dp[1][0] = sticker[1][0]

		for i := 1; i < m; i++ {

			if i == 1 {
				dp[0][i] = dp[1][i-1] + sticker[0][i]
				dp[1][i] = dp[0][i-1] + sticker[1][i]

			} else {
				dp[0][i] = int(math.Max(float64(dp[1][i-1]), float64(dp[1][i-2]))) + sticker[0][i]
				dp[1][i] = int(math.Max(float64(dp[0][i-1]), float64(dp[0][i-2]))) + sticker[1][i]

			}
		}

		fmt.Fprintln(writer, int(math.Max(float64(dp[0][m-1]), float64(dp[1][m-1]))))
	}

}
