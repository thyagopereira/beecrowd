class Solution:
    def fib(self, n: int) -> int:
      dp = [0,1]

      if n > 1:
        for i in range(2, n + 1):
          fib = dp[i - 2] + dp[i - 1] 
          dp.append(fib)

      return dp[n]
