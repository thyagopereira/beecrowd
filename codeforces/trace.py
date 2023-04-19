n = int(input())
cir = sorted([int(x) for x in input().split(" ")])
r = 0
pi = 3.1415926536

for i in range(0, n):
    r += cir[i] * cir[i] * (1 - i%2 * 2)

print(r * pi)