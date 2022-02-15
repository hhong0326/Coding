/**
 * @param {number[][]} triangle
 * @return {number}
 */
 var minimumTotal = function(triangle) {
    
    let dp = new Array(triangle.length).fill(0);
    for (let i=0; i<dp.length; i++ ) {
        dp[i] = new Array(triangle[i].length).fill(0);
    }
    
    dp[0][0] = triangle[0][0];
    
    for (let i=1; i<dp.length; i++) {
        dp[i][0] = dp[i-1][0] + triangle[i][0];
        if (dp[i].length > 2 ) {
            for (let j=1; j<dp[i].length-1; j++) {
                dp[i][j] = Math.min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j];
            }
        }
        
        dp[i][dp[i].length-1] = dp[i-1][dp[i-1].length-1] + triangle[i][triangle[i].length-1];
    }
    
    
    return Math.min(...dp[triangle.length-1]);
    
    
};