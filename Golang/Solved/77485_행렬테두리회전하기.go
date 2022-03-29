package solved

func solution(rows int, columns int, queries [][]int) []int {

	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, columns)
		matrix[i][0] = i*columns + 1
		for j := 1; j < columns; j++ {
			matrix[i][j] = matrix[i][j-1] + 1
		}

	}

	mins := []int{}
	for i := range queries {

		x1, y1, x2, y2 := queries[i][0]-1, queries[i][1]-1, queries[i][2]-1, queries[i][3]-1
		temp := matrix[x1][y1]
		min := matrix[x1][y1]

		i, j := x1, y1

		for i < x2 {
			matrix[i][y1] = matrix[i+1][y1]
			if min > matrix[i][y1] {
				min = matrix[i][y1]
			}
			i++
		}
		for j < y2 {

			matrix[x2][j] = matrix[x2][j+1]
			if min > matrix[x2][j] {
				min = matrix[x2][j]
			}
			j++
		}
		for i > x1 {

			matrix[i][y2] = matrix[i-1][y2]
			if min > matrix[i][y2] {
				min = matrix[i][y2]
			}
			i--
		}
		for j > y1 {

			matrix[x1][j] = matrix[x1][j-1]
			if min > matrix[x1][j] {
				min = matrix[x1][j]
			}
			j--
		}
		matrix[x1][y1+1] = temp
		if min > matrix[x1][y1+1] {
			min = matrix[x1][y1+1]
		}
		mins = append(mins, min)

	}

	return mins
}
