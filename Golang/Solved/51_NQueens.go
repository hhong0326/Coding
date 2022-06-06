package main

import (
	"math"
	"strings"
)

func solveNQueens(n int) [][]string {

	res := [][]string{}

	queen(0, n, 0, [][]string{}, &res)
	return res

}

func queen(level int, n int, pos int, board [][]string, res *[][]string) {

	if !check(level, board, pos) {
		return
	} else if level == n {
		ans := make([]string, len(board))

		for i, str := range board {
			ans[i] = strings.Join(str, "")
		}
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
		queen(level+1, n, i, append(board, str), res)
	}
}

func check(level int, board [][]string, pos int) bool {
	if len(board) <= 1 {
		return true
	}

	for i := 0; i < len(board); i++ {
		if i != level-1 && board[i][pos] == "Q" { // 동일 선상
			return false
		}
		for j := 0; j < len(board[i]); j++ { // 대각선
			if i != level-1 && board[i][j] == "Q" && math.Abs(float64(pos-j)/float64((level-1)-i)) == 1.0 {
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

// 52_NQueens.go Best Solution

// func totalNQueens(n int) int {

//     ans := 0
//     board := Construct(n)

//     var backtrack func(r int)
//     backtrack = func(r int) {
//         if r == n {
//             ans++
//             return
//         }

//         for c := 0; c < n; c++ {
//             if board.Can(r,c) {
//                 board.Set(r,c)
//                 backtrack(r + 1)
//                 board.Reset(r,c)
//             }
//         }
//     }

//     backtrack(0)
//     return ans
// }

// type Board struct {
//     Boards [][]byte
//     Row []bool
//     Col []bool
//     PosD []bool
//     NegD []bool
// }

// func Construct(N int) *Board {
//     board := make([][]byte, N)
//     for i := range board {
//         board[i] = make([]byte, N)
//     }

//     for i := range board {
//         for j := range board[i] {
//             board[i][j] = '.'
//         }
//     }

//     return &Board {
//         Boards: board,
//         Row: make([]bool, N),
//         Col: make([]bool, N),
//         PosD: make([]bool, 2*N),
//         NegD: make([]bool, 2*N),
//     }
// }

// func (b *Board) Set(i, j int) {
//     b.Boards[i][j], b.Row[i], b.Col[j], b.PosD[i+j], b.NegD[i-j+len(b.Boards)]       = 'Q', true, true, true, true
// }

// func (b *Board) Reset(i, j int) {
//     b.Boards[i][j], b.Row[i], b.Col[j], b.PosD[i+j], b.NegD[i-j+len(b.Boards)]       = '.', false, false, false, false
// }

// func (b *Board) Can(i, j int) bool {
//     return !b.Row[i] && !b.Col[j] && !b.PosD[i+j] && !b.NegD[i-j+len(b.Boards)]
// }
