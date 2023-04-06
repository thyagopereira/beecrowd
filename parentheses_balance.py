pair = { "(" :")" , "[" : "]"}

n = int(input())
while n > 0 :
    stack = []
    ok = True
    s = input()

    for l in s:
        if l == "[" or l == "(":
            stack.append(l)
        elif len(stack) > 0:
            if stack[-1] == "(" and l == pair["("]:
                stack.pop()
            elif stack[-1] == "[" and l == pair["["]:
                stack.pop()
        else:
            ok = False
            

    if len(stack) > 0:
        ok = False
    
    if ok:
        print("Yes")
    else:
        print("No")

    n -= 1

