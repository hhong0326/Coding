package main

import (
	"math"
	"strings"
)

func solveNQueens(n int) [][]string {

	res := [][]string{}

	queen(0, n, 0, []string{}, &res)
	return res

}

func queen(level int, n int, pos int, board []string, res *[][]string) {

	if !check(level, board, pos) {
		return
	} else if level == n {
		ans := make([]string, len(board))
		copy(ans, board)
		(*res) = append(*res, ans)
		return
	}
	str := make([]string, n)
	for i := 0; i < n; i++ {
		str[i] = "Q"
		for j := 0; j < n; j++ {
			if i != j {
				str[j] = "."
			}
		}
		queen(level+1, n, i, append(board, strings.Join(str, "")), res)
	}
}

func check(level int, board []string, pos int) bool {
	if len(board) <= 1 {
		return true
	}

	for i := 0; i < len(board); i++ {
		if i != level-1 && string(board[i][pos]) == "Q" { // 동일 선상
			return false
		}
		for j := 0; j < len(board[i]); j++ { // 대각선
			if i != level-1 && string(board[i][j]) == "Q" && math.Abs(float64(pos-j)/float64((level-1)-i)) == 1.0 {
				return false

			}

		}

	}

	return true
}

// other's solution

// func solveNQueens(n int) [][]string {
// 	chess := make([][]byte, n)
// 	for i := range chess {
// 		chess[i] = make([]byte, n)
// 		for j := range chess[i] {
// 			chess[i][j] = '.'
// 		}
// 	}
// 	var res [][]string
// 	solve(&res, &chess, 0)

// 	return res
// }

// func solve(res *[][]string, chess *[][]byte, row int) {
// 	if row == len(*chess) {
// 		*res = append(*res, construct(chess))
// 		return
// 	}

// 	for col := 0; col < len(*chess); col++ {
// 		if isValid(chess, row, col) {
// 			(*chess)[row][col] = 'Q'
// 			solve(res, chess, row+1)
// 			(*chess)[row][col] = '.'
// 		}
// 	}
// }

// func isValid(chess *[][]byte, row, col int) bool {
// 	for i := 0; i < row; i++ {
// 		if (*chess)[i][col] == 'Q' {
// 			return false
// 		}
// 	}

// 	for i, j := row-1, col+1; i >= 0 && j < len(*chess); i, j = i-1, j+1 {
// 		if (*chess)[i][j] == 'Q' {
// 			return false
// 		}
// 	}

// 	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
// 		if (*chess)[i][j] == 'Q' {
// 			return false
// 		}
// 	}

// 	return true
// }

// func construct(chess *[][]byte) []string {
// 	var path []string
// 	for i := 0; i < len(*chess); i++ {
// 		path = append(path, string((*chess)[i]))
// 	}

// 	return path
// }
