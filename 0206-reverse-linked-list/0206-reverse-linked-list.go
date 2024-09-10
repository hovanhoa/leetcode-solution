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

    cur, tail := head, head.Next
    for tail != nil {
        cur = tail
        tail = tail.Next
        cur.Next = head
        if head.Next == cur {
            head.Next = nil
        }

        head = cur
    }

    return cur
}