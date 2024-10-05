/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil {
        return nil
    }

    slow, fast := head, head
    cur := slow
    for fast != nil && fast.Next != nil {
        cur = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    cur.Next = slow.Next

    return head
}