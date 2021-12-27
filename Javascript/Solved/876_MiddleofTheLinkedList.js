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

// Linked List, Two Pointers

var middleNode = function(head) {
    
    var p1 = head;
    var p2 = head;
    var count = 0;
    
    while (p1 != null) {
        p1 = p1.next;
        count++;
    }
    
    count = Math.floor(count/2);
    
    while (count > 0) {
        p2 = p2.next;
        count--;
       
    }
    
    return p2;
};