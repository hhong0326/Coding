/**
 * @param {number[]} nums
 * @return {number}
 */

// DP

var rob = function(nums) {
    
    if (nums.length <= 3) {
        return Math.max(...nums);
    }
    
    var dps = function(s, e) {
        var dp = new Array(s + (e-s) + 1);
        
        for (var i=s; i<=e; i++) {
            dp[i] = Math.max((dp[i-1] ? dp[i-1] : 0), (dp[i-2] ? dp[i-2] : 0) + nums[i]); 
        }
        
        return dp[e];
    }
    
    var dp1 = dps(0, nums.length - 2, nums); // except last index
    var dp2 = dps(1, nums.length - 1, nums);
    
    return Math.max(dp1, dp2);
};