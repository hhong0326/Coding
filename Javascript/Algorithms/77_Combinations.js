/**
 * @param {number} n
 * @param {number} k
 * @return {number[][]}
 */
 var combine = function(n, k) {
        
    let res = [];
    backtrack([], 1);
    return res;
  
    function backtrack(set, num) {
        if (set.length === k) {
            res.push([...set]);
            return;
        }
    
        for (let i = num; i <= n; i++) {
            set.push(i);
            backtrack(set, i + 1);
            set.pop();
        }
    }
    
//     var arr = [];
    
//     for (let i=1; i<=n; i++) {
//         arr.push(i);
//     }
    
//     const result = getCombi(arr, k);
    
//     return result;
};

function getCombi (arr, selectNumber) {
    var results = [];
    if(selectNumber === 1){
        return arr.map((ele) => [ele]);
    }
    arr.forEach((fixed, index, origin) => {
        var rest = origin.slice(index+1);
        var combis = getCombi(rest, selectNumber - 1);
        var attached = combis.map((combi) => [fixed, ...combi]);
        results.push(...attached);
    });
    return results;
}
