class Solution:
    def romanToInt(self, s: str) -> int:
      roman = {"I":1, "V":5, "X": 10, "L":50, "C": 100, "D": 500,    "M": 1000}

      result = 0
      i = 0 
      while (i < len(s) -1 ):
  
        if roman[s[i]] < roman[s[i + 1]]: 
          result += - roman[s[i]] 
        else:
          result += roman[s[i]]

        i += 1
        # Problema:  Só vai até o penultimo, se o ultimo for caso else. naõ vai contabilizar

      
      result += roman[s[-1]]
      
      return result
