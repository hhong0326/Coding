package main

func numMatchingSubseq(s string, words []string) int {

	cnt := 0
	for _, word := range words {
		n := len(word)
		if n > len(s) {
			continue
		} else if n == len(s) || n == 0 {
			if word == s {
				cnt++
			}
		} else {
			i := 0
			for j := 0; j < len(s); j++ {
				if s[j] == word[i] {
					i++
				}
				if i == n {
					cnt++
					break
				}
			}
		}
	}
	return cnt
}
