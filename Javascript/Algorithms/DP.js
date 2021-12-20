// 70. Climbing Stairs by LeetCode

/**
 * @param {number} n
 * @return {number}
 */

// Bottom Up
// var climbStairs = function(n) {
    
//     var dp = new Array(n+1).fill(0);
    
//     a = 1;
//     b = 2;
    
//     var res = a + b;
    
//     for (var i=3; i<=n; i++) {
//         dp[i] = dp[i-1] + dp[i-2];
//     }
    
        
//     return dp[n];
// };


// Top down - Memorization

var climbStairs = function(n) {
    let memo = new Array(50).fill(0);
    
    memo[1] = 1;
    memo[2] = 2;
    
    return recStairs(n, memo);
};

const recStairs = (n, memo) => {
    if (n <= 2) return n;
    
    if (memo[n] === 0) {
        memo[n] = recStairs(n - 1, memo) + recStairs(n - 2, memo);
    }
    
    return memo[n];
}