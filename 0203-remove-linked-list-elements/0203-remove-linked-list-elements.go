/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    dummy := ListNode{Next: head}

    prev, cur := &dummy, head
    for cur != nil {
        fmt.Println(cur.Val)
        if cur.Val == val {
            prev.Next = cur.Next
        } else {
            prev = cur
        }

        cur = cur.Next
    }

    return dummy.Next
}