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

	arr := make([][]int, n)
	var r, g, b int
	dp := make([][]int, n+1)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &r, &g, &b)
		arr[i] = append(arr[i], r, g, b)
		dp[i] = make([]int, 3)
	}

	dp[0][0] = arr[0][0]
	dp[0][1] = arr[0][1]
	dp[0][2] = arr[0][2]

	for i := 1; i < n; i++ {

		dp[i][0] = int(math.Min(float64(dp[i-1][1]), float64(dp[i-1][2]))) + arr[i][0]
		dp[i][1] = int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][2]))) + arr[i][1]
		dp[i][2] = int(math.Min(float64(dp[i-1][0]), float64(dp[i-1][1]))) + arr[i][2]

	}

	fmt.Println(int(math.Min(math.Min(float64(dp[n-1][0]), float64(dp[n-1][1])), float64(dp[n-1][2]))))
}
