/**
 * @param {number[]} nums
 * @return {string[]}
 */
 var summaryRanges = function(nums) {
    
    let res = [];

    for(let i=0; i<nums.length; i++){
        
        let arr = [nums[i]];

        while (nums[i+1] === nums[i]+1) i++; 

        if(arr[0] !== nums[i]) {
            arr.push(nums[i])
        }

        res.push(arr);
    }

    return res.map(arr => arr.join("->"));
};