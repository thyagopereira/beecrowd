n = int(input())
resp = ""

while n >0 and n%7 != 0:
    resp = "4" + resp
    n -= 4

while n > 0 and n%7 == 0:
    resp += "7"
    n -= 7
     
if n == 0:
    print(resp)
else:
    print(-1)

