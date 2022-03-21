/**
 * @param {number[][]} mat
 * @param {number} r
 * @param {number} c
 * @return {number[][]}
 */
 var matrixReshape = function(mat, r, c) {
    
    if (r == 0 || c == 0 || r*c != mat.length*mat[0].length) return mat

    // data set 
    let temp = new Array(mat.length * mat[0].length).fill(0)
    let idx = 0
    for (let i = 0; i < mat.length; i++){
        for (let j = 0; j < mat[0].length; j++){
            temp[idx++] = mat[i][j]
        }
    }
    
    let res = new Array(r).fill(0).map(item => new Array(c).fill(0));
    
    idx = 0
    for(let i = 0; i < r; i++){
        for (let j = 0; j < c; j++){
            res[i][j] = temp[idx++]
            if (res[i][j] == undefined) return mat;
        }
    }
    
    return res
};