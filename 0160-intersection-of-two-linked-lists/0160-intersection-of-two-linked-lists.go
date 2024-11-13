/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    mapA := map[*ListNode]bool{}
    for headA != nil {
        mapA[headA] = true
        headA = headA.Next
    }

    for headB != nil {
        if mapA[headB] {
            return headB
        }

        headB = headB.Next
    }

    return nil
}