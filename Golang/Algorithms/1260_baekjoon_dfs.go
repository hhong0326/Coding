package main

func bfs(lines [][]int) {

}

func main() {

	// reader := bufio.NewReader(os.Stdin)
	// writer := bufio.NewWriter(os.Stdout)

	// var testCaseLength int
	// fmt.Fscanln(reader, &testCaseLength)

	// num1, num2 선언
	// var num1, num2 int
	var lines [][]int

	lines = append(lines, []int{1, 2}, []int{1, 3}, []int{1, 4}, []int{2, 4}, []int{3, 4})

	visited := make([][]bool, 5)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, 5)
		for j := 0; j < len(visited[0]); j++ {
			visited[i][j] = false
		}
	}

	var q [][]int
	q = append(q, lines[0])

	for j := 0; j < len(lines); j++ {
	}

	for len(q) > 0 {

		l := len(q)

		for i := 0; i < l; i++ {
			node := q[0]
			visited[node[0]][node[1]] = true
			visited[node[1]][node[0]] = true
			q = q[1:]

			if visited[node[0]][node[1]] || visited[node[1]][node[0]] {
				continue
			}
			if node[0] == lines[i][0] {
				q = append(q, lines[i])
			}
		}
	}

	// for문을 testCase 길이만큼 돌면서 num1, num2를 입력받고 합을 출력
	// for idx := 0; idx < testCaseLength; idx++ {
	// 	if idx == 0 {
	// 		continue
	// 	}

	// 	fmt.Fscanln(reader, &num1, &num2)
	// 	lines = append(lines, []int{num1, num2})
	// }

	// fmt.Fprintln(writer, num1+num2)

	// writer.Flush()
}
