/**
 * // Definition for a Node.
 * function Node(val, left, right, next) {
 *    this.val = val === undefined ? null : val;
 *    this.left = left === undefined ? null : left;
 *    this.right = right === undefined ? null : right;
 *    this.next = next === undefined ? null : next;
 * };
 */

/**
 * @param {Node} root
 * @return {Node}
 */
var connect = function(root) {
    
    if (root === null) return root;
    
    var q = [];
    q.push(root);
    
    while(q.length) {
        var len = q.length;
        
        for (var i=0; i<len; i++) {
            
            var node = q.shift();
            
            if (i != len -1) {
                node.next = q[0];
            } else {
                node.next = null;
            }
            
            if (node.left != null) {
                q.push(node.left)
            }

            if (node.right != null) {
                q.push(node.right);
            }
        }
        
        
    }
    
    return root;
};