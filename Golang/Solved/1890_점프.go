package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	graph [][]string
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		bytes, _, _ := reader.ReadLine() // string 입력 시
		line := string(bytes)
		// fmt.Fscanln(reader)
		v := strings.Split(line, " ")
		graph = append(graph, v)
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == "0" {
				continue
			}
			jump, _ := strconv.Atoi(graph[i][j])

			if i+jump < n {
				dp[i+jump][j] += dp[i][j]
			}
			if j+jump < n {
				dp[i][j+jump] += dp[i][j]
			}
		}
	}

	fmt.Fprintln(writer, dp[n-1][n-1])
}
