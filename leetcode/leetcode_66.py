class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
      i = len(digits) - 1
      ov = 1
      result = []

      while i >= 0:
        if digits[i] + ov == 10:
          result.insert(0,0)
          ov = 1
        else:
          result.insert(0,digits[i] + ov)
          ov = 0

        i -= 1
      
      if ov > 0:
        result.insert(0,ov)
      return result
          




        

