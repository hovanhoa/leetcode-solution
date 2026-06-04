/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    cur := head
    var prev *ListNode
    for cur != nil {
        cur.Next, cur, prev = prev, cur.Next, cur
    }

    return prev
}