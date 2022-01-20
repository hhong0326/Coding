package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var arr [][]int
	for i := 0; i < n; i++ {
		bytes, _, _ := reader.ReadLine() // string 입력 시
		line := string(bytes)
		v := strings.Split(line, " ")

		var a []int
		for _, vv := range v {
			aa, _ := strconv.Atoi(vv)
			a = append(a, aa)
		}

		arr = append(arr, a)

	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = arr[0][0]

	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {

			if j == 0 {
				dp[i][j] = dp[i-1][j] + arr[i][j]
			} else if j == len(dp[i])-1 {
				dp[i][j] = dp[i-1][j-1] + arr[i][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j-1]), float64(dp[i-1][j]))) + arr[i][j]
			}
		}
	}
	max := dp[n-1][0]

	for _, v := range dp[n-1] {
		if max < v {
			max = v
		}
	}

	fmt.Fprintln(writer, max)
}
