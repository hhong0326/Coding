package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	deque  []int
	reader *bufio.Reader
	writer *bufio.Writer
)

func main() {

	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	for n > 0 {

		bytes, _, _ := reader.ReadLine() // string 입력 시
		line := string(bytes)
		// fmt.Fscanln(reader)
		arr := strings.Split(line, " ")

		switch arr[0] {
		case "push_back":
			a, _ := strconv.Atoi(arr[1])
			pushBack(a)
		case "push_front":
			a, _ := strconv.Atoi(arr[1])
			pushFront(a)
		case "front":
			front()
		case "back":
			back()
		case "size":
			size()
		case "empty":
			answer := empty()
			fmt.Fprintln(writer, answer)
		case "pop_front":
			popFront()
		case "pop_back":
			popBack()
		}

		n--
	}

}

func pushFront(x int) {

	deque = append([]int{x}, deque...)
}

func pushBack(x int) {
	deque = append(deque, x)
}

func popFront() {
	if empty() == 1 {
		fmt.Fprintln(writer, -1)
		return
	}
	v := deque[0]
	deque = deque[1:]

	fmt.Fprintln(writer, v)
}

func popBack() {
	if empty() == 1 {
		fmt.Fprintln(writer, -1)
		return
	}
	v := deque[len(deque)-1]
	deque = deque[:len(deque)-1]

	fmt.Fprintln(writer, v)
}

func size() {

	fmt.Fprintln(writer, len(deque))
}

func empty() int {
	if len(deque) == 0 {
		return 1
	}
	return 0
}

func front() {
	if empty() == 1 {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, deque[0])
}

func back() {
	if empty() == 1 {
		fmt.Fprintln(writer, -1)
		return
	}

	fmt.Fprintln(writer, deque[len(deque)-1])
}
