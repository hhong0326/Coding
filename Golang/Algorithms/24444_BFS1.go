package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n, m, R int
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &R)

	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		arr[i] = append(arr[i], a, b)
	}

	fmt.Println(arr)

}
