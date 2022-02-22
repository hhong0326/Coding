/**
 * @param {number[]} height
 * @return {number}
 */
 var maxArea = function(height) {
    
    let area = 0;
    let i = 0, j = height.length-1;
   
    while (i < j){
        area = Math.max(area, Math.min(height[i], height[j]) * (j - i));
        if (height[i] < height[j]) {
            i++;
        } else {
            j--;
        }
    }
    
    return area;
};