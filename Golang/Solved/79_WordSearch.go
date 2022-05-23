package main

// var (
//     pos = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
// )

// func exist(board [][]byte, word string) bool {

//     if len(board) == 1 {
//         if len(board[0]) == 1 {
//             if len(word) == 1 {
//                 if board[0][0] == word[0] {
//                     return true
//                 }
//             }
//         }
//     }

//     for i:=0; i<len(board); i++ {
//         for j:=0; j<len(board[i]); j++ {
//             if backtrack(board, i, j, word) {
//                 return true
//             }
//         }
//     }

//     return false

// }

// func backtrack(board [][]byte, i, j int, word string) bool {

//     if len(word) < 1 {
//         return true
//     }

//     if board[i][j] != word[0] {
//         return false
//     }

//     board[i][j] = '#'
//     for _, v := range pos {
//         dx := i + v[0]
//         dy := j + v[1]

//         if dx < 0 || dy < 0 || dx >= len(board) || dy >= len(board[i]) {
//             continue
//         }

//         if backtrack(board, dx, dy, word[1:]) {
//             return true
//         }

//     }
//     board[i][j] = word[0]

//     return false
// }

func exist(board [][]byte, word string) bool {

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			if backtrack(board, word, x, y) {
				return true
			}
		}
	}

	return false
}

func backtrack(board [][]byte, word string, x, y int) bool {
	var b byte

	if len(word) < 1 {
		return true
	}
	if y < 0 || y > len(board)-1 {
		return false
	}
	if x < 0 || x > len(board[0])-1 {
		return false
	}
	if board[y][x] != word[0] {
		return false
	}

	board[y][x] = b

	if backtrack(board, word[1:], x, y-1) {
		return true
	}
	if backtrack(board, word[1:], x+1, y) {
		return true
	}
	if backtrack(board, word[1:], x, y+1) {
		return true
	}
	if backtrack(board, word[1:], x-1, y) {
		return true
	}

	board[y][x] = word[0]

	return false
}
