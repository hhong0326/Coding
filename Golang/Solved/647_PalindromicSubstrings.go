package coding

// DP
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	b := []byte(s)
	palindromic := 0
	for i := 0; i < len(b); i++ {
		for j := 0; j <= i; j++ {
			if i-j > 1 {
				dp[j][i] = dp[j+1][i-1] && b[j] == b[i]
			} else {
				dp[j][i] = b[j] == b[i]
			}
			if dp[j][i] {
				palindromic++
			}
		}
	}
	return palindromic
}

// first

// func countSubstrings(s string) int {

//     count := 0;
//     for i:=0; i<len(s); i++ {
//         for j:=i; j<len(s); j++ {
//             if isPalindromic(s[i:j+1]) {
//                 count++
//             }
//         }
//     }

//     return count
// }

// func isPalindromic(ss string) bool{

//     for i, j := 0, len(ss)-1; i<len(ss); i, j = i+1, j-1 {
//         if ss[i] != ss[j] {
//             return false
//         }
//     }

//     return true
// }
