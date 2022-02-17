package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	visited []int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, R int
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &R)

	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		arr[i] = append(arr[i], a, b)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1]
	})

	visited = make([]int, 100001)

	bfs(arr, R)

	for i := 1; i <= n; i++ {
		fmt.Fprintln(w, visited[i])
	}
}

func bfs(arr [][]int, R int) {

	visited[R] = R
	count := 2
	q := []int{R}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, v := range arr {
			if v[0] == node {
				if visited[v[1]] == 0 {
					visited[v[1]] = count
					count++

					q = append(q, v[1])
				}
			}
		}
	}
}
