package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	graph [][]string
	n     int
	pos   = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	fmt.Fscanln(reader, &n)

	var str string
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &str)
		v := strings.Split(str, "")
		graph = append(graph, v)
	}

	var countArr []int

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == "1" {

				cnt := bfs([][]int{{i, j}})
				countArr = append(countArr, cnt)
			}
		}
	}

	fmt.Fprintln(writer, len(countArr))

	sort.Slice(countArr, func(i, j int) bool { return countArr[i] < countArr[j] })
	for _, v := range countArr {
		fmt.Fprintln(writer, v)
	}
}

func bfs(q [][]int) int {

	var cnt int
	for len(q) > 0 {

		a := q[0]
		q = q[1:]

		if a[0] < 0 || a[1] < 0 || a[0] >= n || a[1] >= n || graph[a[0]][a[1]] != "1" {
			continue
		}

		graph[a[0]][a[1]] = "2"
		cnt++

		for i := 0; i < len(pos); i++ {
			q = append(q, []int{a[0] + pos[i][0], a[1] + pos[i][1]})
		}
	}

	return cnt
}
