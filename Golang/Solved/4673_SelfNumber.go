package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	// d[n] = n + if n >= 10 && n/10 + n%10

	m := make(map[int]int, 10000)

	for i := 1; i <= 10000; i++ {

		var sum int
		a := i

		sum += a
		for a >= 10 {
			a /= 10
			sum += a
		}
		a %= 10
		sum += a

		m[sum] = i

		if v, ok := m[i]; ok {
			fmt.Println(v)
		}
	}

}
