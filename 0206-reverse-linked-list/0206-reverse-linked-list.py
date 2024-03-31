# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if not head:
            return None
        
        iterNode = head
        head = head.next
        iterNode.next = None

        while head:
            temp = head
            head = head.next
            temp.next = iterNode
            iterNode = temp

        return iterNode