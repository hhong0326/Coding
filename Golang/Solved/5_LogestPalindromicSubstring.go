package main

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		max = maxPalindrome(s, i, i, max)
		max = maxPalindrome(s, i, i+1, max)
	}
	return max
}

func maxPalindrome(s string, i, j int, max string) string {
	var sub string
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}

	if len(max) < len(sub) {
		return sub
	}
	return max
}

// func longestPalindrome(s string) string {

//     m := make(map[string]bool)
//     dp := make([][]string, len(s))

//     for i:=0; i<len(s); i++ {
//         for j:=i+1; j<=len(s); j++ {
//             if _, ok := m[s[i:j]]; ok {
//                 continue
//             }
//             if check(s[i:j]) {
//                 // if max < len(s[i:j]) {
//                 //     max = len(s[i:j])
//                 // }
//                 dp[len(s[i:j])-1] = append(dp[len(s[i:j])-1], s[i:j])
//                 m[s[i:j]] = true
//             }
//         }
//     }

//     for i:=len(dp)-1; i>=0; i-- {
//         if len(dp[i]) > 0 {
//             return dp[i][0]
//         }
//     }

//     return ""
// }

// func check(str string) bool {
//     for i:=0; i<len(str)/2; i++ {
//         if str[i] != str[len(str)-1-i] {
//             return false
//         }
//     }

//     return true
// }
