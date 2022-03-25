/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
 var sumOfLeftLeaves = function(root) {
    
    if (!root) return 0;
    
    let answer = 0;
    
    let s = [root];
    
    while (s.length > 0 ) {
        
        const node = s.pop();
        
        if (node.left && isLastLeft(node.left)) {
            answer += node.left.val;
        }
        if (node.left) {
            s.push(node.left);
        }
        if (node.right) {
            s.push(node.right);
        }
    }
    
    return answer;
};

var isLastLeft = function(node) {
    
    return !node.left && !node.right
}