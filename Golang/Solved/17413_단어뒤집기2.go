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

	bytes, _, _ := r.ReadLine() // string 입력 시
	str := string(bytes)

	words := strings.Split(str, "")

	// tag check
	// < 나오면 > 나오는 때 까지 문자열을 저장한다

	var i int
	var s int
	for i < len(words) {

		if words[i] == "<" {
			i += 1
			for words[i] != ">" {
				i += 1
			}
			i += 1
		} else if words[i] != " " && words[i] != "<" && words[i] != ">" {
			s = i
			for i < len(words) && words[i] != " " && words[i] != "<" && words[i] != ">" {
				i += 1
			}
			temp := words[s:i]
			arr := reverse(temp)
			for j := s; j < i; j++ {
				words[j] = arr[j-s]

			}
		} else if words[i] == " " {
			i += 1
		}
	}

	fmt.Fprintln(w, strings.Join(words, ""))
}

func reverse(arr []string) []string {

	for i, j := 0, len(arr)-1; i < len(arr)/2; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
