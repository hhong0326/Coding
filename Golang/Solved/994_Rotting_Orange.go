package coding

// BFS
func orangesRotting(grid [][]int) int {

	var total, rotted int
	n, m := len(grid), len(grid[0])
	q := make([][2]int, 0)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 || grid[i][j] == 2 {
				total++
				if grid[i][j] == 2 {
					q = append(q, [2]int{i, j})
				}
			}
		}
	}

	if total == 0 {
		return 0
	}

	var timer int
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			rotted++
			node := q[0]
			q = q[1:]
			for _, d := range directions {
				x, y := node[0]+d[0], node[1]+d[1]
				if x < 0 || y < 0 || x == n || y == m || grid[x][y] != 1 { // 0 || 2
					continue
				}
				grid[x][y] = 2
				q = append(q, [2]int{x, y})
			}
		}
		timer++
	}

	if rotted == total {
		return timer - 1
	}
	return -1
}
