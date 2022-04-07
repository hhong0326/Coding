package main

func isAnagram(s string, t string) bool {

	//	   m := make(map[rune]int, 26)
	//     m2 := make(map[rune]int, 26)

	//     for _, v := range s {
	//         m[v-'a']++
	//     }

	//     for _, v := range t {
	//         m2[v-'a']++
	//     }

	//     for i:=0; i<26; i++ {
	//         if m[rune(i)] != m2[rune(i)] {
	//             return false
	//         }
	//     }

	m := make(map[int]int, 26)

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		m[int(s[i]-'a')]++
		m[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if m[i] != 0 {
			return false
		}
	}

	return true
}
