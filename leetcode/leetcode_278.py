class Solution:
    def firstBadVersion(self, n: int) -> int:
      b = 0

      while b <= n :
        p = int( (b + n) / 2)
        if isBadVersion(p):
          if not isBadVersion(p -1):
            return p
          else:
            n = p - 1
        else:
          b = p + 1  
