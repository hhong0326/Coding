/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
 var addTwoNumbers = function(l1, l2) {
    let head = new ListNode();
    let curr = head;
    let carry = 0;
    while (l1 || l2) {
        let v1 = l1 ? l1.val : 0;
        let v2 = l2 ? l2.val : 0;
        let sum = v1 + v2 + carry
        carry = sum > 9 ? 1 : 0;
        if (l1) {
            l1.val = sum % 10;
            curr.next = l1;
        } else {
            l2.val = sum % 10;
            curr.next = l2;
        }
        curr = curr.next;
        if (l1) l1 = l1.next;
        if (l2) l2 = l2.next;
    }
    if (carry > 0) {
        curr.next = new ListNode(1);
    }
    return head.next;
};