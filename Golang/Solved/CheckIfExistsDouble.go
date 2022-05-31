package main

func checkIfExist(arr []int) bool {

	// hashTable

	m := make(map[int]bool)

	for _, n := range arr {
		if n%2 == 0 {
			if _, ok := m[2*n]; ok {
				return true
			}

			if _, ok := m[n/2]; ok {
				return true
			}
		} else {
			if _, ok := m[2*n]; ok {
				return true
			}
		}
		m[n] = true
	}

	return false

	//     for i:=0; i<len(arr); i++ {
	//         for j:=0; j<len(arr); j++ {
	//             if i == j {
	//                 continue
	//             }
	//             if arr[i] * 2 == arr[j] {

	//                 return true
	//             }
	//         }
	//     }

	//     return false

}
