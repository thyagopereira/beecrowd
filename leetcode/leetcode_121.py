class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        
        b,s = 0,0
        _max = 0
        while s < len(prices) :
            diff = prices[s] - prices[b]
            if diff >= 0: 
                if diff > _max :
                    _max = diff
            else:
                b = s

            s += 1

        return _max
            
