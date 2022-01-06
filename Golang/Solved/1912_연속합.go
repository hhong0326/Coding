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
	var arr []int

	for i := 0; i < testCaseLength; i++ {
		fmt.Fscan(reader, &n)
		arr = append(arr, n)
	}

	dp := make([]int, testCaseLength+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = -1001
	}

	answer := -1001

	for i := 1; i <= testCaseLength; i++ {
		dp[i] = int(math.Max(float64(dp[i-1]+arr[i-1]), float64(arr[i-1])))
		answer = int(math.Max(float64(answer), float64(dp[i])))
	}

	fmt.Fprintln(writer, answer)
	writer.Flush()
}
