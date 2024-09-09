/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    var slow, fast, prev *ListNode = head, head, nil
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }

    prev.Next = slow.Next

    return head
}