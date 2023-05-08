# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def removeNthFromEnd(self, head, n):
        """
        :type head: ListNode
        :type n: int
        :rtype: ListNode
        """
        a = ListNode(val = 0, next = head)    
        l, r = a, a

        for i in range(n):
            r = r.next

        while r.next is not None :
            r = r.next
            l = l.next

        # Simplesmente remove
        l.next = l.next.next
        
        return a.next