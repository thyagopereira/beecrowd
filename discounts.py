n = int(input())
prices = sorted([int (x) for x in input().split(" ")])

n_cupons = int(input())
c = [int (x) for x in input().split(" ")]

for cp in c:
    
    cost = 0

    for i in range(len(prices) -1, -1, -1):
        if i != len(prices) - cp:
            cost += prices[i]

    print(cost)