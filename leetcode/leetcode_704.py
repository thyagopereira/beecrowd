class Solution:
    def search(self, nums: List[int], target: int) -> int:
      b,e = 0, len(nums) -1 

      while e >= b :
        p = int( (e + b) / 2)

        if nums[p] == target:
          return p
        elif target < nums[p]:
          e = p - 1 
        else:
          b = p + 1

      return -1
