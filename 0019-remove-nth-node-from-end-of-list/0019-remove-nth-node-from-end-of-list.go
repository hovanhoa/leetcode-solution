/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{Next: head}
    left := dummy
    right := head

    for n > 0 && right != nil {
        right = right.Next
        n -= 1
    }

    for right != nil {
        left = left.Next
        right = right.Next
    }

    left.Next = left.Next.Next
    return dummy.Next
}