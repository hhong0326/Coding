package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	q := make([]int, n)

	for i := 0; i < n; i++ {
		q[i] = i + 1
	}

	for len(q) != 1 {
		q = q[1:]
		temp := q[0]
		q = q[1:]
		q = append(q, temp)
	}

	fmt.Fprintln(writer, q[0])
}
