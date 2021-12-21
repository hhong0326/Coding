package solved

// BFS

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image[sr][sc] == newColor {
		return image
	}

	pos := [4][2]int{
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, 0},
	}

	origin := image[sr][sc]

	q := make([][]int, 0)
	q = append(q, []int{sr, sc})

	// bfs - queue

	for len(q) > 0 {
		c := q[0]
		q = q[1:]

		if c[0] < 0 || c[1] < 0 || c[0] >= len(image) || c[1] >= len(image[0]) || image[c[0]][c[1]] != origin {
			continue
		}

		image[c[0]][c[1]] = newColor

		for i := 0; i < len(pos); i++ {
			q = append(q, []int{c[0] + pos[i][0], c[1] + pos[i][1]})
		}
	}

	return image
}
