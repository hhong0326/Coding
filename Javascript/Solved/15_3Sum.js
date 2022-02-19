/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var threeSum = function(nums) {
    
    let results = [];
    
    
    nums.sort((a, b) => a-b);
    
    for (let i=0; i<nums.length-2; i++) {
        if (nums[i] !== nums[i-1]) {
            let s = i+1, e = nums.length-1;
            let sum = 0 - nums[i];
            
            while (s < e) {
                if (nums[s] + nums[e] == sum) {
                    results.push([nums[i], nums[s], nums[e]]);
                    while(s < e && nums[s] == nums[s+1]) s++; // 중복 방지
                    while(e > s && nums[e] == nums[e-1]) e--;
                    
                    s++;
                    e--;
                } else if(nums[s] + nums[e] < sum) {
                    s++;
                } else {
                    e--;
                }
            }
        }
        
    }
    
    return results;
};
