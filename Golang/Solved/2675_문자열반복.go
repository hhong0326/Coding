package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	var c int
	var str string

	for i := 0; i < n; i++ {

		fmt.Fscanf(r, "%d %s\n", &c, &str)
		arr := strings.Split(str, "")

		fmt.Fprintln(w, repeat(c, arr))
	}
}

func repeat(cnt int, arr []string) string {

	var str string
	for _, v := range arr {
		for i := 0; i < cnt; i++ {
			str += v
		}
	}

	return str
}
