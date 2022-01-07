package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	fmt.Fscanln(reader, &n)

	answer := 0

	for {
		if n%5 != 0 {
			if n < 3 {
				answer = -1
				break
			}
			n -= 3
			answer++
		} else {
			break
		}
	}
	if answer != -1 {
		answer += n / 5
	}

	fmt.Fprintln(writer, answer)
	writer.Flush()

}
