/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
    length := 0
    current := head
    for current != nil {
        length++
        current = current.Next
    }

    current = head
    base, extra := length/k, length%k
    res := make([]*ListNode, 0, k)
    for i := 0; i < k; i ++ {
        path := base
        if extra > i {
            path++
        }

        var partHead, partTail *ListNode
        for j := 0; j < path; j++ {
            if partHead == nil {
                partHead, partTail = current, current
            } else {
                partTail.Next = current
                partTail = partTail.Next
            }

            if current != nil {
                current = current.Next
            }
        }

        if partTail != nil {
            partTail.Next = nil
        }

        res = append(res, partHead)
        
    }

    return res
}