package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	arr []int
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	var n, m int

	fmt.Fscanln(r, &n)

	arr = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d ", &arr[i])
	}

	sort.Ints(arr)

	fmt.Fscanln(r, &n)

	finds := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(r, "%d ", &finds[i])
	}

	var answer []int
	for _, v := range finds {

		vv := bs(v)
		if vv == -1 {

			answer = append(answer, 0)
		} else {
			answer = append(answer, vv)

		}

	}

	for _, v := range answer {
		fmt.Fprintf(w, "%d ", v)
	}
}

func bs(target int) int {

	s, e := 0, len(arr)-1

	for s <= e {
		mid := (s + e) / 2

		if arr[mid] == target {

			arr2 := arr[s : e+1]
			return len(arr2)

		} else if arr[mid] < target {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return -1
}
