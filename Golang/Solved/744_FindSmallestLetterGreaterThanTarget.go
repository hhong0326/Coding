package coding

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {

	s, e := 0, len(letters)
	for s < e {
		mid := s + (e-s)/2
		if letters[mid] <= target {
			s = mid + 1
		} else {
			e = mid
		}

		fmt.Println(s)
	}
	return letters[s%len(letters)]
}
