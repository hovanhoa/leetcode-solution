/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    cur, nxt := head, head.Next
    for nxt != nil {
        if cur.Val != nxt.Val {
            cur.Next = nxt
            cur = nxt
        }

        nxt = nxt.Next
    }

    cur.Next = nxt

    return head
}