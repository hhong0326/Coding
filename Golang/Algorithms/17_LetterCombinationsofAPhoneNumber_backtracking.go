package main

// Backtracking
func letterCombinations(digits string) []string {

	if len(digits) == 0 {
		return []string{}
	}

	hashMap := make(map[byte]string, 0)
	hashMap['2'] = "abc"
	hashMap['3'] = "def"
	hashMap['4'] = "ghi"
	hashMap['5'] = "jkl"
	hashMap['6'] = "mno"
	hashMap['7'] = "pqrs"
	hashMap['8'] = "tuv"
	hashMap['9'] = "wxyz"

	res := []string{}
	backtrack(0, "", &digits, &res, hashMap)
	return res

}

func backtrack(idx int, arr string, digits *string, res *[]string, hashMap map[byte]string) {

	if len(*digits) == idx {
		*res = append(*res, arr)
		return
	}

	num := (*digits)[idx]
	list := hashMap[num]

	for _, v := range list {
		// arr += string(v)
		backtrack(idx+1, arr+string(v), digits, res, hashMap)
		// arr = arr[:len(arr)-1]
	}

}
