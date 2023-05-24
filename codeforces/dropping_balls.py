while True:
    k = int(input())
    if k == -1:
        break
    
    while k > 0:
        entry = [int(x) for x in input().split(" ")]
        x, y = entry[0], entry[1]

        ans = 1
        for i in range(x-1):
            if y & 1:
                ans <<= 1
                y = (y + 1) >> 1
            else:
                ans = (ans << 1) + 1
                y >>= 1
        print(ans)
        k -= 1