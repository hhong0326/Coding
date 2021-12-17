/**
 * @param {character[][]} matrix
 * @return {number}
 */

// DP 

var maximalSquare = function(matrix) {
    
    if (matrix.length == 0) return 0;
    
    // multiple Array DP
//     var dp = new Array(matrix.length+1).fill(0).map(item => item = new Array(matrix[0].length+1).fill(0));
    
//     var max = 0;
//     for (var i=1; i<=matrix.length; i++) {
//         for (var j=1; j<=matrix[0].length; j++) {
//             if (matrix[i-1][j-1] == '1') {
//                 dp[i][j] = Math.min(Math.min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1;
//                 max = Math.max(max, dp[i][j]);
                
//             }
            
//         }
//     }
    
    // single Array DP
    // columns
    
    var dp2 = new Array(matrix[0].length+1).fill(0);
    var max = 0, prev = 0;
    
    for (var i=1; i<=matrix.length; i++) {
        for (var j=1; j<=matrix[0].length; j++) {
            var temp = dp2[j];
            if (matrix[i-1][j-1] == '1') {
                dp2[j] = Math.min(Math.min(dp2[j-1], prev), dp2[j]) + 1;
                max = Math.max(max, dp2[j]);
                
            } else {
                dp2[j] = 0;
            }
            
            prev = temp;
        }
    }
    
    return max * max;
};