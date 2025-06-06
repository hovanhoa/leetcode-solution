/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    

    var prev *ListNode
    for slow != nil {
        tmp := slow.Next
        slow.Next = prev
        prev = slow
        slow = tmp
    }

    l, r := head, prev
    for r != nil {
        if l.Val != r.Val {
            return false
        }

        l, r = l.Next, r.Next
    }

    return true
}