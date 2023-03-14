class Solution:
  def lengthOfLongestSubstring(self, s: str) -> int:

    rep = {}
    aux = ""
    i,j,c, _max = 0, 0, 0, 0

    while i < len(s) and j < len(s):
      
      if s[j] in rep:
        i += 1
        j = i
        c = 1
        rep = {}
      else:
        c += 1
      
      rep[s[j]] = None
      _max = max(c, _max)
      j += 1

    return _max
      
