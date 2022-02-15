/**
 * @param {number[]} nums
 * @return {number}
 */
 var singleNumber = function(nums) {
    
    //     let obj = {};
        
    //     for (let i=0; i<nums.length; i++){
    //         if (obj[nums[i]]) {
    //             obj[nums[i]]++;
    //         } else {
    //             obj[nums[i]] = 1;
    //         }
        
    //     }
        
    //     let result = Object.keys(obj).filter(item => obj[item] == 1);
        
    //     return result[0]
        
        let missing = 0;
        
        for (let i=0; i<nums.length; i++) {
            const val = nums[i]
            missing = missing ^ val // XOR 비트연산자 다르면 1, 같으면 0
        }
        
        return missing;
    };