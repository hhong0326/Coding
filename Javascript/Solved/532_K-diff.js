/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
 var findPairs = function(nums, k) {
    
    let cnt = 0;
    let obj = {};
    nums.sort((a, b) => a-b);
    //Brute force

//     for (let i=0; i<nums.length; i++) {
//         for (let j=i+1; j<nums.length; j++) {
//             if (k == Math.abs(nums[i]-nums[j])) {
                
//                 if (obj[nums[i]] == nums[j]) {
//                     continue;
//                 } else {
//                     obj[nums[i]] = nums[j]
//                     cnt++;
//                 }
                
//             }
//         }
//     }
    
    // Two pointer
    let s = 0, e = 1;
    
    while (s <= nums.length-1 && e <= nums.length-1) {
        
        if (s == e || nums[e] - nums[s] < k){
            e++
        } else if (nums[e] - nums[s] > k){
            s++;
        } else {
            s++
            cnt++
            while(s <= nums.length-1 && nums[s] == nums[s-1])
                s++
        }
    }
    
    
    
    return cnt;
};