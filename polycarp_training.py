n = int(input())
q = sorted([int(x) for x in input().split(" ")])

d = 1
count = 0 
for i in range(len(q)):
    if q[i] >= d:
        d += 1
        count += 1

print(count)