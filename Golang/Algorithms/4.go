package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	count = 0
	sum   = 0
	n     int
	k     int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &n, &k)

	arr := make([]int, k)
	combination(arr, 0, k, 0)

	fmt.Fprintln(w, count)
}

func combination(arr []int, idx, kk, target int) {
	if kk == 0 {
		for _, v := range arr {
			sum += v
		}

	} else if target == n {
		return
	} else {

		arr[idx] = target

		combination(arr, idx+1, kk-1, target+1)
		combination(arr, idx, kk, target+1)
	}
}
