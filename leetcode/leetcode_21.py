class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
      if list1 == None and list2 == None:
        return None
      else:
        resp = ListNode()
        i, j, k = list1, list2, resp

      

        while i != None and j != None:
          
          if i.val <= j.val:
            k.val = i.val
            i = i.next
          else:
            k.val = j.val
            j = j.next
          
          k.next = ListNode()
          k = k.next

        if i != None:
          while i != None:
            k.val = i.val
            i = i.next

            if i != None:  
              k.next = ListNode()
              k = k.next

        elif j != None:
          while j != None:
            k.val = j.val
            j = j.next

            if j != None:
              k.next = ListNode()
              k = k.next

        return resp

