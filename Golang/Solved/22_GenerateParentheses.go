package main

// // dp
// func generateParenthesis(n int) []string {

//     // left := "("
//     // right := ")"

//     dp := make([]map[string]bool, 17)

//     dp[0] = make(map[string]bool, 1)
//     dp[0][""] = true
//     dp[1] = make(map[string]bool, 2)
//     dp[1]["("] = true
//     dp[1][")"] = true

//     for i:=2; i<len(dp); i++ {
//         dp[i] = make(map[string]bool, 0)

//         for j:=1; j<i; j++ {
//             for k := range dp[j] {
//                 for kk := range dp[i-j] {

//                     s := k + kk
//                     if _, ok := dp[i][s]; !ok {
//                         dp[i][k + kk] = true
//                     }

//                 }
//             }

//         }

//     }

//     ans := []string{}
//     for k := range dp[n*2] {
//         if isRight(k) {
//             ans = append(ans, k)
//         }
//     }

//     return ans
// }

// func isRight(str string) bool {

//     var s []string

//     for _, v := range str {
//         if string(v) == "(" {
//             s = append(s, string(v))
//         } else {
//             if len(s) == 0 {
//                 return false
//             } else {
//                 s = s[:len(s)-1]
//             }
//         }
//     }
//     if len(s) > 0 {
//         return false;
//     }

//     return true
// }

// backtracking
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(n, n, &res, "")
	return res
}

func backtrack(left, right int, res *[]string, cur string) {
	if left == 0 && right == 0 {
		*res = append(*res, cur)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		backtrack(left-1, right, res, cur+"(")
	}

	if right > 0 {
		backtrack(left, right-1, res, cur+")")
	}
}
