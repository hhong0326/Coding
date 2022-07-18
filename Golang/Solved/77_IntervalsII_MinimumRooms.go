package main

import (
	"fmt"
	"math"
	"sort"
)

func solution(timeTables [][]int) {

	sTimetables := []int{}
	eTimetables := []int{}

	for i := 0; i < len(timeTables); i++ {
		sTimetables = append(sTimetables, timeTables[i][0])
		eTimetables = append(eTimetables, timeTables[i][1])
	}

	// O(nlogn)
	sort.Ints(sTimetables)
	sort.Ints(eTimetables)

	fmt.Println(sTimetables)
	fmt.Println(eTimetables)

	s, e := 0, 0
	rooms, maxRooms := 0, 0
	for s < len(sTimetables) || e < len(eTimetables) {
		if s >= len(sTimetables) {
			break
		}

		if sTimetables[s] < eTimetables[e] {
			rooms++
			s++
		} else {
			rooms--
			e++
		}

		maxRooms = int(math.Max(float64(maxRooms), float64(rooms)))
	}

	fmt.Println(maxRooms)
}

func main() {

	solution([][]int{{30, 75}, {0, 50}, {60, 150}})
	solution([][]int{{0, 1}, {0, 2}, {0, 3}, {2, 3}})

}
