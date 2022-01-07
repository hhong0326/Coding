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

	var stairs int
	fmt.Fscanln(reader, &stairs)

	var score []int

	var n int
	for i := 0; i < stairs; i++ {
		fmt.Fscanln(reader, &n)
		score = append(score, n)
	}

	// score = append(score, 10, 20, 15, 25, 10, 20)

	dp := make([]int, stairs+1)
	for i := 0; i < stairs; i++ {

		if i == 0 {
			dp[i] = score[i]
		} else if i == 1 {

			dp[i] = score[0] + score[i]
		} else if i == 2 {
			dp[i] = int(math.Max(float64(score[0]), float64(score[1]))) + score[i]

		} else {
			dp[i] = int(math.Max(float64(dp[i-3])+float64(score[i-1]), float64(dp[i-2]))) + score[i]
		}

	}

	fmt.Fprintln(writer, dp[stairs-1])
	writer.Flush()
}
