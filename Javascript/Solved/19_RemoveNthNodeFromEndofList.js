/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} n
 * @return {ListNode}
 */
var removeNthFromEnd = function(head, n) {
    
    var p1 = head;
    var count = 0;
    
    while(p1 !== null){
        p1 = p1.next;
        count++;
    }
    
    if (count == 1) {
        return null;
    }
    
    count = count - n - 1;
    p1 = head;
    
    if (count == -1) {
        return p1.next;
    }
    
    var idx = 0;
    
    while(idx != count && p1 != null) {
        p1 = p1.next;
        idx++;
    }
    
    if (p1.next != undefined) {
        p1.next = p1.next.next;
    } else {
        p1 = null;
    }
    
    
    return head;
};