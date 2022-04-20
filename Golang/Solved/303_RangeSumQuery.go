package main

// Immutable
type NumArray struct {
	arr []int
}

// normal sum
// func Constructor(nums []int) NumArray {
//     return NumArray{nums}
// }

// func (this *NumArray) SumRange(left int, right int) int {
//     var sum int
//     for _, v := range this.arr[left:right+1] {
//         sum += v
//     }

//     return sum
// }

// accumulate
func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{[]int{}}
	}
	tmp := make([]int, len(nums))
	tmp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		tmp[i] = tmp[i-1] + nums[i]
	}
	return NumArray{arr: tmp}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.arr[j]
	}
	return this.arr[j] - this.arr[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
