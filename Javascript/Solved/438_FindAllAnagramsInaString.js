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

// new solution guide, using charCodeAt

var findAnagrams = function (s, p) {
    const chars = new Array(26).fill(0), res = [];
    
    for(let i = 0; i < p.length; i++) {
        chars[p.charCodeAt(i) - 97]--;
    }
    
    main:
    for(let i = 0; i < s.length; i++){
        chars[s.charCodeAt(i) - 97]++;

        if(i < p.length - 1) continue;
        if(i > p.length - 1) chars[s.charCodeAt(i - p.length) - 97]--;

        for(let j = 0; j < 26; j++){
            if(chars[j]) continue main;
        }
        
        res.push(i - p.length + 1);
    }
    
    return res;
};