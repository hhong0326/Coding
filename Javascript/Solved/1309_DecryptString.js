/**
 * @param {string} s
 * @return {string}
 */
 var freqAlphabets = function(s) {
    
    let answer = "";
    let stack = "";
    
    for (let i=s.length-1; i>=0; i--) {
        if (s[i] == '#') {
            stack += s.slice(i-2, i);
            i-=2
            
        } else {
            stack += s[i];   
        }
        
        answer = String.fromCharCode(Number(stack)+96) + answer;
        stack = "";
    }
    
    return answer;
};