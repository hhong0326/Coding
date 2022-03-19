/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node|null} root
 * @return {number[]}
 */
 var preorder = function(root) {
    
    if (root == null) return []
    if (root.children.length == 0 ) return [root.val]
    
    let s = [];
    let answer = [];
    
    s.push(root);
    
    while (s.length > 0) {
        node = s.pop();
        answer.push(node.val);
        
        if (node.children.length == 0) {
            continue;
        }
        
        for (let i=node.children.length-1; i>=0; i--) {
            s.push(node.children[i]);
        }
    }

    return answer;
};
