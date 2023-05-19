n = int(input())
dp = [0] * (n + 2)
dp[0] = 1

for i in range(1, n + 1):
    for j in range(1, 7):
        if j <= i:
            dp[i] = (dp[i] + dp[i -j]) % (10**9 + 7) 

print(dp[n])