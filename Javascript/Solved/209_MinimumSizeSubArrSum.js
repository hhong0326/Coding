/**
 * @param {number} target
 * @param {number[]} nums
 * @return {number}
 */
 var minSubArrayLen = function(target, nums) {
    
    var min = Number.POSITIVE_INFINITY;
    var sum = 0;
    var start = 0;
    for(var i in nums) {
        sum += nums[i];
        while(sum >= target) {
            min = Math.min(min, i-start+1);
            sum -= nums[start++];
        }
    }
    return min === Number.POSITIVE_INFINITY ? 0 : min;
};
 
