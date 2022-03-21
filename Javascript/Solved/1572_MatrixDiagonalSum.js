/**
 * @param {number[][]} mat
 * @return {number}
 */
 var diagonalSum = function(mat) {
    
    let sum = 0;
    
    if (mat.length % 2 == 0) {
        for (let i=0; i<mat.length; i++) {
            sum += mat[i][i]
            sum += mat[i][mat.length-1-i];
        }
        
    } else {
        for (let i=0; i<mat.length; i++) {
            
            if (i !== mat.length-1-i) {
                sum += mat[i][i]
                sum += mat[i][mat.length-1-i];
           
            } else {
                 sum += mat[i][i]
            }
        }
    }
    
    return sum
};