package main

func findDisappearedNumbers(nums []int) []int {

	m := make(map[int]bool, 0)

	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}

	res := []int{}
	for i := 1; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			res = append(res, i)
		}
	}

	return res

	// using in-place

	//     i := 0
	//     n := len(nums)

	//     for i < n {
	//         v := nums[i]-1
	//         if nums[i] != nums[v]{
	//             nums[i],nums[v] = nums[v],nums[i]
	//         }else{
	//             i++
	//         }
	//     }

	//     result := []int{}
	//     for i,v := range nums {
	//         if v != i + 1 {
	//             result = append(result, i+1)
	//         }
	//     }

	//     return result

	// using arr index

	//     s := make([]bool, len(nums)+1)
	//     s[0] = true

	//     for _, num := range nums {
	//         s[num] = true
	//     }

	//     res := make([]int, 0)
	//     for num, val := range s {
	//         if !val {
	//             res = append(res, num)
	//         }
	//     }
	//     return res
}
