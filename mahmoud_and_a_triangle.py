n = int(input())
retas = sorted([int(x) for x in input().split(" ")])
t = False

for i in range(0, len(retas) -2 ):
    a = retas[i]
    b = retas[i+1]
    c = retas[i+2]

    if ((a + b > c) and (c + b > a) and (a + c > b)):
        t = True

        break

if t:
    print("YES")
else:
    print("NO")
