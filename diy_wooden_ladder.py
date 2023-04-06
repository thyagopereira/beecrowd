t = int(input())
sizes = {}
answer = []

for i in range(t):
    n = int(input())
    sizes[i] = sorted([int(x) for x in input().split(" ")])

for i in range(t):
    b = min(sizes[i][-1], sizes[i][-2]) -1
    s = 0
    for j in range(0, len(sizes[i]) - 2):
        if s < b:
            s += 1

    answer.append(s)


for a in answer:
    print(a)

