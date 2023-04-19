class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:

      if head:
        prev = head
        curr = head.next 
    
        while curr:
          if curr.val == prev.val:
            prev.next = curr.next
            curr = curr.next
          else:
            prev = curr
            curr = curr.next
            
      return head
