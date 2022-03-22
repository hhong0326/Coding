package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	L, C    int
	results []string
	list    []string
	w       *bufio.Writer
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanf(r, "%d %d\n", &L, &C)

	list = make([]string, C)
	for i := 0; i < C; i++ {
		fmt.Fscanf(r, "%s", &list[i])
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	var vlist []string
	var clist []string

	for i := 0; i < len(list); i++ {

		switch list[i] {
		case "a", "e", "i", "o", "u":
			vlist = append(vlist, list[i])
		default:
			clist = append(clist, list[i])
		}
	}

	bt([]string{}, 0)

}

func bt(set []string, num int) {
	if len(set) == L {

		newSet := strings.Join(set, "")
		if !check(newSet) {
			return
		}

		fmt.Fprintln(w, newSet)
		return
	}

	for i := num; i < len(list); i++ {
		set = append(set, list[i])
		bt(set, i+1)
		set = set[:len(set)-1]
	}
}

func check(str string) bool {

	vCnt := 0
	cCnt := 0

	for i := 0; i < len(str); i++ {
		switch string(str[i]) {
		case "a", "e", "i", "o", "u":
			vCnt++
		default:
			cCnt++
		}
	}

	if vCnt >= 1 && cCnt >= 2 {
		return true
	}

	return false
}
