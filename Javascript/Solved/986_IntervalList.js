/**
 * @param {number[][]} firstList
 * @param {number[][]} secondList
 * @return {number[][]}
 */
 var intervalIntersection = function(firstList, secondList) {
    
    let s1 = 0, s2 = 0;
    let answer = [];
    
    while (s1 < firstList.length && s2 < secondList.length) {
        
        let min = Math.max(firstList[s1][0], secondList[s2][0]);
        let max = Math.min(firstList[s1][1], secondList[s2][1]);
        
        if (min <= max) {
            answer.push([min, max]);
        }
       
        if (firstList[s1][1] < secondList[s2][1]) {
            s1++;
        } else {
            s2++;
        }
            
    }
    
    return answer;
};