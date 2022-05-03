/**
 * @param {number[]} nums
 * @return {number}
 */
 var numberOfArithmeticSlices = function(nums) {
    
    // at least
    
    if (nums.length < 3) {
        return 0
    }
    var count = 0;
    var curIndex = 0;
    for (var i=1; i<nums.length-1; i++) {
        if (nums[i] - nums[i-1] === nums[i+1] - nums[i]) {
            count +=  i - curIndex;
        }
        else {
            curIndex = i;
        }
    }

    return count;
};