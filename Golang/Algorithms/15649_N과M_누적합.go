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

	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	//1부터 N까지 자연수 중에서 중복없이 M개 짜리 수열 (오름차순)
	visited := make([]bool, n+1)
	answer := make([]int, n+1)

	dfs(writer, visited, answer, 0, n, m)
}

func dfs(writer *bufio.Writer, visited []bool, answer []int, idx, n, m int) {
	if idx == m {
		for i := 0; i < m; i++ {
			fmt.Fprintf(writer, "%d ", answer[i])
		}
		fmt.Fprint(writer, "\n")
		return
	}
	for i := 0; i < n; i++ {
		if visited[i] == true {
			continue
		}
		answer[idx] = i + 1
		visited[i] = true
		dfs(writer, visited, answer, idx+1, n, m)
		visited[i] = false
	}
}
