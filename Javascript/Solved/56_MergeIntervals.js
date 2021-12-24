/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */

// Array, Sorting

var merge = function(intervals) {
    
    var newArr = [];
    intervals.sort((a, b) => a[0] - b[0]);
    
    var l = intervals[0][0], r = intervals[0][1];
    
    for (var i=1; i<intervals.length; i++){
      
        if (r >= intervals[i][0]) {
            r = Math.max(r, intervals[i][1]);
            continue;
        }
        
        newArr.push([l, r]);
        l = intervals[i][0]; // next
        r = intervals[i][1];
    
    }
    
    newArr.push([l, r]);
    
    return newArr;
};