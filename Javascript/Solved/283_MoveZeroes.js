/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */

// Array, Two Pointers

var moveZeroes = function(nums) {
    
//     var i = 0;
//     while (i < nums.length) {
//         if (nums[i] === 0) {
//             nums.push(nums.splice(i, 1));
//             continue;
//         }
        
//         i++;
//     }
    
    // optimal
        
    var j = 0, temp;
    
    for (var i=0; i<nums.length; i++) {
        if (nums[i] !== 0) {
            temp = nums[i];
            nums[i] = nums[j];
            nums[j] = temp; // in front of zero
            j++;
            console.log(nums)
        }
    }
}
        