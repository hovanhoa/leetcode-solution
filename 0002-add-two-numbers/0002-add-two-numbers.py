# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        resNode = iterNode = ListNode()
        n = 0
        while l1 or l2:
            if not l2:
                m = l1.val
                l1 = l1.next
            elif not l1:
                m = l2.val
                l2 = l2.next
            else:
                m = l1.val + l2.val
                l1, l2 = l1.next, l2.next

            m, n = m + n, 0
            if m < 10:
                iterNode.next = ListNode(m)
            else:
                iterNode.next = ListNode(m%10)
                n = m//10
                iterNode.next.next = ListNode(n)
            iterNode = iterNode.next

        return resNode.next


