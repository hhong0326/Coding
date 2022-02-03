package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n)

	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	fmt.Fscanln(reader, &m)

	finds := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &finds[i])
	}

	sort.Ints(arr)

	for i := 0; i < m; i++ {

		s, e := 0, len(arr)-1

		for s <= e {
			mid := (s + e) / 2

			if arr[mid] < finds[i] {
				s = mid + 1
			} else {
				e = mid - 1
			}
		}

		if s == len(arr) || arr[s] != finds[i] {
			fmt.Fprintln(writer, 0)
		}
		fmt.Fprintln(writer, 1)
	}
}
