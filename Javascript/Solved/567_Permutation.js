/**
 * @param {string} s1
 * @param {string} s2
 * @return {boolean}
 */
 var checkInclusion = function(s1, s2) {
    let c = 0, obj = {};
    
	for (let i = 0; i < s1.length; i++) {
	  obj[s1[i]] = obj[s1[i]] ? obj[s1[i]] + 1 : obj[s1[i]] = 1
	}
    
    let co = {...obj};
    
    for(let i = 0; i < s2.length; i++) {
        if(c === s1.length) break;

        if(co[s2[i]] && co[s2[i]]!== 0) {
            co[s2[i]]--;
            c++;
        } else {
            co = {...obj};
            i = i - c;
            c = 0;
        }
    }
    
    return  c === s1.length ? true : false;
};