nx = input().split(" ")
coins = [int(c) for c in input().split(" ")]

x = int(nx[1])
dp = [0] * (x + 1)
dp[0] = 1

for i in range(1, x + 1): 
    for c in coins:
        if i - c >= 0:
            dp[i] = (dp[i] + dp[i - c]) % (10** 9 + 7)

print(dp[x])