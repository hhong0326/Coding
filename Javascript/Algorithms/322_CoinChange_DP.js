/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */

// DP

var coinChange = function(coins, amount) {
   
    var max = amount + 1;
    
    var dp = new Array(max).fill(max);
    dp[0] = 0;
    
    // max amount, min coins
    for (var i=1; i<= amount; i++) {
        for (var j=0; j<coins.length; j++) {
            if (coins[j] <= i) {
                dp[i] = Math.min(dp[i], dp[i - coins[j]] + 1);
                console.log(i, dp[i])
            }
        }
    }
    
    return dp[amount] > amount ? -1 : dp[amount];
};