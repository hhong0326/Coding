/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */

 var insertionSortList = function(head) {
    
    if (head == null) return head;
    
    var p = head;
    var pp = head;
    
    while (pp.next != null) {
        while (p.next != null) {
            var temp;
            if (p.val > p.next.val) {
                temp = p.val;
                p.val = p.next.val;
                p.next.val = temp;
            }

            p = p.next;
        }
        
        pp = pp.next;
        p = head;
    }
    
    
    
    return head;
};