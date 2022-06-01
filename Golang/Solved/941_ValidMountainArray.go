package main

func validMountainArray(arr []int) bool {

	// 25ms

	//     if len(arr) <= 1 {
	//         return false
	//     }

	//     up, down := false, false

	//     for i:=0; i<len(arr)-1; i++ {
	//         if arr[i] < arr[i+1] {
	//             up = true
	//             if down {
	//                 return false
	//             }
	//         } else if arr[i] == arr[i+1] {
	//             return false
	//         } else {
	//             if !up {
	//                 return false
	//             } else {
	//                 down = true
	//             }
	//         }
	//     }

	//     if up && !down {
	//         return false
	//     }

	//     return true

	// other's solution
	// 44ms

	n := len(arr)
	s, e := 0, n-1
	for s+1 < n && arr[s] < arr[s+1] {
		s++
	}

	for e-1 > 0 && arr[e] < arr[e-1] {
		e--
	}
	return s > 0 && e < n-1 && s == e
}
