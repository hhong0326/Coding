package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var s string
	fmt.Fscanln(reader, &s)

	arr := make([]int, 26)
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}

	for i, v := range s {
		key := int(math.Abs(float64(v - 97)))
		// fmt.Println(key)
		if arr[key] != -1 {
			continue
		}
		arr[key] = i
	}

	// // 문자열 hash
	for _, v := range arr {
		fmt.Fprintf(writer, "%d ", v)
	}

}
