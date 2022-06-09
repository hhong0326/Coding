package main

func twoSum(numbers []int, target int) []int {

	s, e := 0, len(numbers)-1

	for numbers[s]+numbers[e] != target {

		sum := numbers[s] + numbers[e]

		if sum < target {
			s++
		} else if sum > target {
			e--
		} else {
			return []int{s + 1, e + 1}
		}
	}

	return []int{}

}
