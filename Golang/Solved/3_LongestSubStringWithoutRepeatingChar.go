package main

func lengthOfLongestSubstring(s string) int {

	start := 0
	max := 0
	m := make(map[byte]int)

	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok && m[s[i]] >= start {
			start = m[s[i]] + 1
		}

		if max < i-start+1 {
			max = i - start + 1
		}
		m[s[i]] = i
	}

	return max
}
