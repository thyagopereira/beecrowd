class Solution:
  def isValid(self, s: str) -> bool:
    op = ['(', '[', '{']
    cl = [')','}', ']']

    pair = {'(': ')', '[':']',  '{': '}'}
    stack = []
    
    if len(s) <= 1:
      return False

    for c in s:
      if c in op:
        stack.append(c)
      elif c in cl:
        if len(stack) > 0:
          tmp = stack[-1]
          stack.pop()

          if pair[tmp] != c:
            return False
        else:
          return False

    return True and len(stack) == 0


