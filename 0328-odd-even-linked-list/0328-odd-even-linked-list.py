# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def oddEvenList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        odd = oddHead = ListNode()
        even = evenHead = ListNode()
        isOdd = True

        while head:
            if isOdd:
                odd.next = head
                odd = head
            else:
                even.next = head
                even = head
            
            head = head.next
            isOdd = not isOdd
        
        even.next = None
        odd.next = evenHead.next
        
        return oddHead.next