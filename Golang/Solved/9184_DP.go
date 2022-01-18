package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	memo   [][][]int
	writer *bufio.Writer
)

// memo

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	for {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}

		memo = make([][][]int, 21)
		for i := 0; i < 21; i++ {
			memo[i] = make([][]int, 21)
			for j := 0; j < 21; j++ {
				memo[i][j] = make([]int, 21)
			}
		}
		answer := w(a, b, c)

		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, answer)
	}
}

func w(a, b, c int) int {

	if memo[a][b][c] != 0 {
		return memo[a][b][c]
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}
	if a > 20 || b > 20 || c > 20 {
		return w(20, 20, 20)
	}
	if a < b && b < c {
		memo[a][b][c] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
	} else {
		memo[a][b][c] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	}

	return memo[a][b][c]
}
