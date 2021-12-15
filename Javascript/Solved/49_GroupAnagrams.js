/**
 * @param {string[]} strs
 * @return {string[][]}
 */

// Hash Table

 var groupAnagrams = function(strs) {
    
    // using sort
//     var obj = {};
//     strs.forEach(s => {
//         var newS = s.slice();
//         newS = newS.split('').sort().join('');
//         if (!obj[newS]) {
//             obj[newS] = [s];
//         } else {
//             obj[newS].push(s);
//         }
//     });
    
    
//     return Object.values(obj);
    
    // using table
    var res = {};
    strs.forEach(str => {
        var count = new Array(26).fill(0);
        
        for (var c of str) count[c.charCodeAt()-97]++; // a 
        var key = count.join('#');
        res[key] ? res[key].push(str) : res[key] = [str];
    })
    
    return Object.values(res);
};