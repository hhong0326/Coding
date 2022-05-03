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
    
    if (root == null) {
		return null
	}

    var p = root;
    var q = [p];
    
    
    while (q.length > 0) {
        
        var l = q.length;
        
        var temp = [];
        
        for (var i=0; i<l; i++) {
            var node = q.shift();
            
            if (node.left) {
                q.push(node.left);
            }
            if (node.right) {
                q.push(node.right);
            }
            
            
            temp.push(node)
        }
        
        
        for (var i=0; i<temp.length-1; i++) {
            temp[i].next = temp[i+1];
        }
         
    }
    
    return root
    
    
};