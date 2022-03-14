package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	pos = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m int
	fmt.Fscanf(r, "%d %d\n", &n, &m)

	var str string

	maze := make([][]string, n)
	visited := make([][][]bool, 50)
	for i := 0; i < 50; i++ {
		visited[i] = make([][]bool, 50)

		for j := 0; j < 50; j++ {
			visited[i][j] = make([]bool, 1<<6)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &str)

		maze[i] = strings.Split(str, "")
	}

	start := func(i, j int) int {

		q := [][]int{
			{i, j},
			{0, 0},
		}
		visited[i][j][0] = true
		maze[i][j] = "."

		for len(q) > 0 {

			p := q[0]
			pp := q[1]

			count := pp[0]
			key := pp[1]

			if maze[p[0]][p[1]] == "1" {
				return count
			}

			q = q[2:]

			for i := 0; i < len(pos); i++ {

				di := p[0] + pos[i][0]
				dj := p[1] + pos[i][1]

				if di >= 0 && di < n && dj >= 0 && dj < m {

					if !visited[di][dj][key] {
						if maze[di][dj] == "." || maze[di][dj] == "1" {

							visited[di][dj][key] = true
							q = append(q, []int{di, dj}, []int{count + 1, key})

						} else if "a" <= maze[di][dj] && maze[di][dj] <= "f" {

							temp := key | (1 << int([]rune(maze[di][dj])[0]-'a'))
							visited[di][dj][temp] = true
							q = append(q, []int{di, dj}, []int{count + 1, temp})

						} else if "A" <= maze[di][dj] && maze[di][dj] <= "F" {
							v := key & (1 << int([]rune(maze[di][dj])[0]-'A'))
							if v != 0 {
								visited[di][dj][key] = true
								q = append(q, []int{di, dj}, []int{count + 1, key})
							}

						}
					}
				}
			}

		}

		return -1
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if maze[i][j] == "0" {
				c := start(i, j)
				fmt.Fprintln(w, c)
			}
		}
	}
}
