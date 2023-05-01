while(True):
    n = int(input())

    if n == 0 :
        break

    street = [int(x) for x in input().split(" ")]
    ans, sum = 0, 0

    for i in street:
        sum += i
        ans += abs(sum)

    print(ans)


#Toma TLE
