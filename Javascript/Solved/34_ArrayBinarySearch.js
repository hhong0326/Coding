/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
 var searchRange = function(nums, target) {
    let s = 0, e = nums.length-1;
    
    let answer = [];
    
    while (s <= e) {
        
        let mid = Math.floor((s + e)/2);
        
        if (nums[mid] === target) {
            while(true) {
                if (nums[s] !== target) {
                    s++;
                } else if (nums[e] !== target) {
                    e--;
                } else {
                    answer.push(s, e);
                    return answer.length == 0 ? [-1, -1] : answer;
                }
            }
            
            
        } else if (nums[mid] < target) {
            s = mid + 1;
        } else {
            e = mid -1;
        }
    }
    
    return [-1, -1]
};