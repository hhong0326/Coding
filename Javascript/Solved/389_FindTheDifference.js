/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 */
 var findTheDifference = function(s, t) {
    
    
    //     let m = new Map();
    //     let answer = "";
        
    //     for (let i=0; i<s.length; i++){
            
    //         let v = m.get(s[i]);
    //         if (v == undefined || !v) {
    //             m.set(s[i], 1);
    //         } else {
    //             m.set(s[i], m.get(s[i])+1)
    //         }
    //     }
        
    //     for (let i=0; i<t.length; i++) {
    //         if (!m.get(t[i])) {
    //             answer += t[i];
    //         } else {
    //             m.set(t[i], m.get(t[i])-1)
    //         }
    //     }
        
    //     for (let v in m) {
    //         if (v > 0) {
    //             answer += v
    //         }
    //     }
    
    //     return answer;
    
    let sSum = 0, tSum = 0, i = 0;
    
    while(i < s.length) {
        sSum += s.charCodeAt(i);
        tSum += t.charCodeAt(i);
        i++;
    }
    tSum += t.charCodeAt(i);
    return String.fromCharCode(tSum - sSum);
};