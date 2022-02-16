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
 var swapPairs = function(head) {
    
    if (head == null) return null;
    let newHead = new ListNode();
    let p = newHead;
    
    let prev = null, cur = head, len = 0;
    while (cur !== null) {
        let next = cur.next;
        cur.next = prev;
        prev = cur;
        cur = next;
        len++;
        if (len % 2 == 0) { 
            p.next = prev;
            p = p.next.next; 
        }
    }
    
    if (len % 2 == 1) {
        p.next = prev;
        p = p.next; 
    }
    p.next = null;
    
    return newHead.next;
};