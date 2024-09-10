/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    oddHead, oddTail := head, head
    evenHead, evenTail := head.Next, head.Next
    for oddTail != nil && evenTail != nil {

        oddTail.Next = evenTail.Next
        if oddTail.Next != nil {
            oddTail = oddTail.Next
        }

        evenTail.Next = oddTail.Next
        evenTail = evenTail.Next
    }    

    oddTail.Next = evenHead

    return oddHead
}