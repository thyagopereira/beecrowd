class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:      
      lp = "" 
      for i in range(len(strs[0])):
        p = strs[0][0:i + 1]

        _diff = True
        for s in strs:
          if s[0:i + 1] != p:
            return lp

        if _diff:
          lp = p

      return lp
        
