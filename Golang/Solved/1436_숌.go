package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdin)

	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	var count int
	var num = 0
	for {

		num++
		str := strconv.Itoa(num)
		if strings.Contains(str, "666") {
			count++
			if count == n {
				fmt.Println(num)
				break
			}
		}
	}

	// for count != n {
	// 	num++
	// 	temp := num

	// 	for temp != 0 {
	// 		if temp%1000 == 666 {
	// 			count++
	// 			break
	// 		} else {
	// 			temp /= 10
	// 		}
	// 	}
	// }

	// fmt.Println(num)
}
