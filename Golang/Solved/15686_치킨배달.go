package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int

	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	var graph [][]string
	mmap := make([][]int, n)
	for i := 0; i < n; i++ {
		bytes, _, _ := reader.ReadLine() // string 입력 시
		line := string(bytes)
		// fmt.Fscanln(reader)
		v := strings.Split(line, " ")
		graph = append(graph, v)
		mmap[i] = make([]int, n)
	}

	fmt.Println(graph)

	var hs [][]int
	var cs [][]int

	// 각 집에서 최소 치킨집 거리 계산
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if graph[i][j] == "2" {
				cs = append(cs, []int{i, j})
			}
			if graph[i][j] == "1" {
				hs = append(hs, []int{i, j})
			}
		}
	}

	var dist []int
	var x, y int
	for i := 0; i < len(hs); i++ {
		min := 50
		for j := 0; j < len(cs); j++ {
			a := int(math.Abs(float64(cs[j][0])-float64(hs[i][0])) + math.Abs(float64(cs[j][1])-float64(hs[i][1])))
			if min > a {
				min = a
				x = cs[j][0]
				y = cs[j][1]
			}
		}

		mmap[x][y] += 1
		dist = append(dist, min)
	}

	fmt.Println(mmap)

	if len(cs) == m {
		var sum int
		for _, v := range dist {
			sum += v
		}
		fmt.Fprintln(writer, sum)
		return
	}

	var sum int
	for _, v := range dist {
		sum += v
	}
	fmt.Fprintln(writer, sum)
	// 한 치킨집이 가지는 모든 최소 거리의 총합이 가장 큰 치킨집을 지워야함
}
