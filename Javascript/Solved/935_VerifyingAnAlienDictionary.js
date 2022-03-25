/**
 * @param {string[]} words
 * @param {string} order
 * @return {boolean}
 */
 var isAlienSorted = function(words, order) {
    
    let obj = {};
    
    for (let i=1; i<=order.length; i++) {
        obj[order[i-1]] = i;
    }
    
    for (let i=0; i<words.length-1; i++) {
        
        for (let j=0; j<words[i].length; j++ ){
            
            if (j >= words[i+1].length) {
                return false;
            }
            
            if (obj[words[i][j]] > obj[words[i+1][j]] ) {
                return false;
            } else if (obj[words[i][j]] == obj[words[i+1][j]] ){
                continue;
            } else {
                break;
            }
        } 
    }
    
    return true
};