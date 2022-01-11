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

	var n, m int
	var arr []int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &m)
		arr = append(arr, m)
	}

	for _, v := range arr {
		zero, one := fibonacci(v)
		fmt.Fprintf(writer, "%d %d\n", zero[v], one[v])

	}
}

func fibonacci(n int) (zero, one []int) {

	zero = []int{1, 0}
	one = []int{0, 1}

	if n < 2 {
		return
	}
	for i := 2; i <= n; i++ {
		zero = append(zero, zero[i-1]+zero[i-2])
		one = append(one, one[i-1]+one[i-2])
	}

	return
}
