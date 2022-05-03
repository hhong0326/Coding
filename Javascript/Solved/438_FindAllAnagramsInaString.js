/**
 * @param {string} s
 * @param {string} p
 * @return {number[]}
 */
 var findAnagrams = function(s, p) {
    
    var answer = [];
    var obj = {};
    for (var i=0; i<p.length; i++) {
        if (obj[p[i]]) {
            obj[p[i]]++;
        } else {
            obj[p[i]] = 1;
        }
    }
    
    var start = 0, end = p.length;
    
    while (start + end <= s.length) {
        
        
        var checkObj = {...obj};
        
        var check = true
        for (var i=start; i<start+end; i++) {
            if (checkObj[s[i]] > 0) {
                checkObj[s[i]]--;
            } else {
                check = false
                break;
            }
        }
        
        if (!check) {
            start++;
            continue;
        } 
        
        var arr = Object.values(checkObj);
        
        for (var i=0; i<arr.length; i++) {
            if (arr[i] > 0) {
                check = false
                break
            }
        }
        
        if (check) {
            answer.push(start);
        } 
        start++;
        
    }
    return answer
    
};