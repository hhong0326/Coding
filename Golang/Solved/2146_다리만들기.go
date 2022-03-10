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

	group := graph

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}

	pos := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}

	dfs := func(i, j, num int) {

		s := [][]int{
			{i, j},
		}

		for len(s) > 0 {
			c := s[len(s)-1]
			s = s[:len(s)-1]

			if c[0] < 0 || c[1] < 0 || c[0] >= n || c[1] >= n || graph[c[0]][c[1]] != "1" {
				continue
			}

			if visited[c[0]][c[1]] {
				continue
			}

			visited[c[0]][c[1]] = true
			group[c[0]][c[1]] = strconv.Itoa(num)

			for k := 0; k < len(pos); k++ {
				di := c[0] + pos[k][0]
				dj := c[1] + pos[k][1]

				s = append(s, []int{di, dj})
			}
		}

	}

	bfs := func(i, j int) int {
		dis := make([][]int, n)
		for i := 0; i < n; i++ {
			dis[i] = make([]int, n)
		}
		land, _ := strconv.Atoi(group[i][j])

		q := [][]int{
			{i, j},
		}
		dis[i][j] = 1

		for len(q) > 0 {

			p := q[0]
			q = q[1:]

			if vv, _ := strconv.Atoi(group[p[0]][p[1]]); land < vv {
				return dis[p[0]][p[1]] - 2
			}

			for k := 0; k < len(pos); k++ {
				di := p[0] + pos[k][0]
				dj := p[1] + pos[k][1]

				if di < 0 || dj < 0 || di >= n || dj >= n {
					continue
				}

				if group[di][dj] == strconv.Itoa(land) {
					continue
				}
				if dis[di][dj] > 0 {
					continue
				}

				dis[di][dj] = dis[p[0]][p[1]] + 1
				q = append(q, []int{di, dj})
			}

		}
		return 200
	}

	number := 2
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == "1" && !visited[i][j] {
				dfs(i, j, number)
				number++
			}
		}
	}

	min := 200
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v, _ := strconv.Atoi(group[i][j]); v > 1 {

				min = int(math.Min(float64(min), float64(bfs(i, j))))
			}
		}
	}

	fmt.Fprintln(w, min)

}
