package main

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	str := words[0] + "#"

	for i := 1; i < len(words); i++ {
		if strings.Contains(str, words[i]+"#") {
			continue
		}
		str += words[i] + "#"
	}
	return len(str)
}
