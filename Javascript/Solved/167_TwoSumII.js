/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */


var twoSum = function(numbers, target) {
    
    
    var s = 0, e = numbers.length-1;
    
    while (s < e) {
        var sum = numbers[s] + numbers[e];
        if (sum === target) {
            return [s + 1, e + 1];
        } else if (sum < target) {
            s++;
        } else {
            e--;
        }
    }
    
    return [-1, -1];
};