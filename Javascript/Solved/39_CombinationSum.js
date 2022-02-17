/**
 * @param {number[]} candidates
 * @param {number} target
 * @return {number[][]}
 */
 var combinationSum = function(candidates, target) {
    
    let sum = 0;
    let results = new Array();
    
    var backtrack = function(path, t, idx) {
        
        if (t < 0) {
            return
        }
        
        if (t == 0) {
            results.push([...path]);
            
            return
        } 
        
        for (let i=idx; i<candidates.length; i++) {
            path.push(candidates[i]);
            backtrack(path, t-candidates[i], i);
            path.pop();
            
        }
    }
    
    backtrack([], target, 0)

    return results
};