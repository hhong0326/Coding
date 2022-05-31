package main

import "math"

func hasAllCodes(s string, k int) bool {

	m := make(map[string]bool)

	//     for i:=0; i<=len(s)-k; i++ {
	//         m[s[i:i+k]] = true
	//     }

	//     res := []string{}
	//     backtrack(0, k, &res, "")

	//     fmt.Println(len(res))

	//     for _, v := range res {
	//         if _, ok := m[v]; !ok {
	//             return false
	//         }
	//     }

	//     return true

	// length solution

	for i := 0; i <= len(s)-k; i++ {
		m[s[i:i+k]] = true
	}

	return len(m) == int(math.Pow(2.0, float64(k)))

}

func backtrack(idx, k int, res *[]string, str string) {

	if idx == k {
		(*res) = append(*res, str)
		return
	}

	backtrack(idx+1, k, res, str+"0")
	backtrack(idx+1, k, res, str+"1")

}
