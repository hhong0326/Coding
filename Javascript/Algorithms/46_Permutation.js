/**
 * @param {number[]} nums
 * @return {number[][]}
 */
 var permute = function(nums) {
    
    const res = [];
    const backtrack = (path) => {
        if(path.length === nums.length){
            res.push(path);
            return;
        }
        nums.forEach(n => {
           if(path.includes(n)){return ;}
            backtrack(path.concat(n));
        });
    };
    backtrack([]);
    return res;
    
    
//     const result = getPermute(nums, nums.length)
//     return result;
};

function getPermute (arr, selectNumber) {
    var results = [];
    if(selectNumber === 1){
        return arr.map((ele) => [ele]);
    }
    arr.forEach((fixed, index, origin) => {
        var rest = [...origin.slice(0, index), ...origin.slice(index + 1)];
        var combis = getPermute(rest, selectNumber - 1);
        var attached = combis.map((combi) => [fixed, ...combi]);
        results.push(...attached);
    });
    return results;
}