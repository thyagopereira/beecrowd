class Solution:
  def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:

    aux_l1, aux_l2, i, b = 0,0,0,10

    while l1 or l2:
      
      if l1:
        aux_l1 += l1.val * (b**i) 
        l1 = l1.next
      if l2: 
        aux_l2 += l2.val * (b**i)
        l2 = l2.next

      i += 1

    r = str(aux_l1 + aux_l2)
    res = ListNode()
    rl = res
    
    for i in range(len(r) -1, -1,-1):

      if i != 0:
        rl.val = int(r[i])
        rl.next = ListNode()
        rl = rl.next
      else:
        rl.val = int(r[i])

    return res




        
