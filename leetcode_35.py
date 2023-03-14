class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
      b,e = 0, len(nums) - 1

      while b <= e:
        p = int((b + e) / 2)
        if nums[p] == target:
          return p
        elif target < nums[p]:
          e = p - 1
        else:
          b = p + 1
      
      if target < nums[p]:
        return p
      else:
        return p + 1 
