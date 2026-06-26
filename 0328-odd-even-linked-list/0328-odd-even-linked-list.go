/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    left, right := &ListNode{
        Next: head,
    }, &ListNode{
        Next: head,
    }
    curLeft, curRight := left, right
    cnt := 0
    for head != nil {
        cnt += 1
        if cnt%2 == 1 {
            curLeft.Next = head
            curLeft = curLeft.Next
        } else {
            curRight.Next = head
            curRight = curRight.Next
        }

        head = head.Next
    }

    curRight.Next = nil
    curLeft.Next = right.Next
    return left.Next
}