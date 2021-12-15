package coding

import "fmt"

func tartgetNumber(numbers []int, target int) int {
	var answer int = 0

	result := getCombinations([]int{1, 2, 3}, 2)
	fmt.Println("in", result)
	return answer
}

func dfsTN() {

}

func getCombinations(array []int, selectNumber int) [][]int {

	var results [][]int
	if selectNumber == 1 {
		return append(results, array)
	}
	for i, v := range array {
		rest := array[i+1:]
		combinations := getCombinations(rest, selectNumber-1)

		attached := []int{v}
		for _, v := range combinations {
			attached = append(attached, v...)
		}
		results = append(results, attached)
	}

	return results
}

func StartTN() {

	result := tartgetNumber([]int{1, 1, 1, 1, 1}, 3)
	fmt.Println(result)
}
