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

	var n int
	fmt.Fscanln(reader, &n)

	var str string
	var arr []string
	var answer string
	for i := 0; i < n; i++ {

		fmt.Fscanln(reader, &str)
		arr = strings.Split(str, "")

		if isCorrect(arr) {
			answer = "YES"
		} else {
			answer = "NO"
		}

		fmt.Fprintln(writer, answer)
	}

}

func isCorrect(arr []string) bool {

	s := []string{}

	for i := 0; i < len(arr); i++ {

		switch arr[i] {
		case "(":
			s = append(s, arr[i])

		case ")":
			if len(s) != 0 {
				s = s[:len(s)-1]
			} else {
				return false
			}
		}
	}

	if len(s) != 0 {
		return false
	}

	return true
}
