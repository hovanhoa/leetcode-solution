
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var beforeSlow *ListNode
	for fast != nil && fast.Next != nil {
		beforeSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	beforeSlow.Next = slow.Next

	return head
}