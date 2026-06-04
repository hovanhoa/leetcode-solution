/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    dummy := &ListNode{
        Next: head,
    }

    if head == nil || left == right {
        return head
    }

    // 1. reach note at position left 
    leftPrev, cur := dummy, head
    for i := 0; i < left-1; i++ {
        leftPrev = cur
        cur = cur.Next
    }

    // 2. reverse from left to right
    var prev *ListNode
    for i := left; i <= right; i++ {
        tmpNext := cur.Next
        cur.Next = prev
        prev, cur = cur, tmpNext
    }

    leftPrev.Next.Next = cur
    leftPrev.Next = prev

    return dummy.Next    
}