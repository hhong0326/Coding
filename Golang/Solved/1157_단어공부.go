package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var str string

	fmt.Fscanln(reader, &str)

	arr := strings.Split(str, "")

	m := make(map[string]int, 0)

	for _, v := range arr {
		vv := strings.ToUpper(v)
		if _, ok := m[vv]; ok {
			m[vv]++
		} else {
			m[vv] = 1
		}
	}

	var lastest string
	var max int
	var count = 1
	for k, v := range m {
		if max <= v {
			if max == v {
				count++
			} else {
				max = v
				lastest = k
				count = 1
			}
		}
	}

	if count > 1 {
		fmt.Fprintln(writer, "?")
		return
	}
	fmt.Fprintln(writer, lastest)
}
