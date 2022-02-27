package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	var str string
	var graph [][]string
	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &str)

		strs := strings.Split(str, "")
		graph = append(graph, strs)
	}

	visited, rgVisited := make([][]bool, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
		rgVisited[i] = make([]bool, n)
	}

	pos := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	var count int
	var rgCount int

	rgdfs := func(i, j int) {

		s := [][]int{
			{i, j},
		}

		for len(s) > 0 {
			c := s[len(s)-1]
			s = s[:len(s)-1]

			if c[0] < 0 || c[1] < 0 || c[0] >= n || c[1] >= n || graph[c[0]][c[1]] == "B" {
				continue
			}
			if rgVisited[c[0]][c[1]] {
				continue
			}

			rgVisited[c[0]][c[1]] = true

			for k := 0; k < len(pos); k++ {
				di := c[0] + pos[k][0]
				dj := c[1] + pos[k][1]

				s = append(s, []int{di, dj})
			}
		}

	}

	dfs := func(i, j int, color string) {

		s := [][]int{
			{i, j},
		}

		for len(s) > 0 {
			c := s[len(s)-1]
			s = s[:len(s)-1]

			if c[0] < 0 || c[1] < 0 || c[0] >= n || c[1] >= n {
				continue
			}

			if color == "R" || color == "G" {
				rgdfs(i, j)
			}
			if visited[c[0]][c[1]] || color != graph[c[0]][c[1]] {
				continue
			}

			visited[c[0]][c[1]] = true

			for k := 0; k < len(pos); k++ {
				di := c[0] + pos[k][0]
				dj := c[1] + pos[k][1]

				s = append(s, []int{di, dj})
			}
		}

	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				if !rgVisited[i][j] {
					rgCount++
				}
				dfs(i, j, graph[i][j])
				count++
			}
		}
	}

	fmt.Fprintln(w, count, rgCount)
}
