package main

func removeDuplicates(nums []int) int {

	//     count := 0
	//     for i:=0; i<len(nums); i++ {

	//         if i == 0 || nums[i] != nums[i - 1] {

	//             count++

	//         }
	//     }

	//     pos := 0

	//     for i:=0; i<len(nums); i++ {
	//         if i==0 || nums[i] != nums[i-1] {
	//             nums[pos] = nums[i]
	//             pos++
	//         }
	//     }

	//     return count

	// prev := nums[0]
	// end := 1
	// for _, v := range nums[1:] {
	//     if v == prev {
	//         continue
	//     }
	//     nums[end] = v
	//     prev = v
	//     end++
	// }
	// return end

	// Check for edge cases.
	if len(nums) == 0 {
		return 0
	}

	// Use the two pointer technique to remove the duplicates in-place.
	// The first element shouldn't be touched; it's already in its correct place.
	writePointer := 1
	// Go through each element in the Array.
	for readPointer := 1; readPointer < len(nums); readPointer++ {
		// If the current element we're reading is *different* to the previous
		// element...
		if nums[readPointer] != nums[readPointer-1] {
			// Copy it into the next position at the front, tracked by writePointer.
			nums[writePointer] = nums[readPointer]
			// And we need to now increment writePointer, because the next element
			// should be written one space over.
			writePointer++
		}
	}

	// This turns out to be the correct length value.
	return writePointer

}
