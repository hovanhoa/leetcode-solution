/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    next := head
    var tail *ListNode
    for head != nil {
        next = head.Next
        head.Next = tail
        tail = head
        head = next
    }

    return tail
}