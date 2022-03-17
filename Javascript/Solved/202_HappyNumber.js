/**
 * @param {number} n
 * @return {boolean}
 */
 var isHappy = function(n) {
    
    
    var ns = String(n);
    
    var result = '';
    
    var s = new Set();
    
    while(result != '1') {
        
        var b = [];
        
        for (var i=0; i<ns.length; i++) {
            b.push(Math.pow(Number(ns[i]), 2));
        }
        
        result = b.reduce((pre, next) => pre+next);
        
        
        if (s.has(result)) {
            return false;
        }
        
        s.add(result);
        ns = String(result);
     
    }

    return true;
};