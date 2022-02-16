package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	answer := bfs(arr, R)

	for _, v := range answer {
		fmt.Fprintln(w, v)
	}
}

func bfs(arr [][]int, R int) []int {

	visited := make([]bool, len(arr)+1)
	answer := make([]int, len(arr))

	visited[R] = true
	answer[0] = R
	count := 1
	q := []int{R}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		for _, v := range arr {
			if v[0] == node {
				if !visited[v[1]] {
					visited[v[1]] = true
					answer[count] = v[1]
					count++

					q = append(q, v[1])
				}
			}
		}
	}

	return answer
}
