class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        d = {}
        for n in nums:
            if n in d.keys():
                return True
            else:
                d[n] = None
        
        return False
