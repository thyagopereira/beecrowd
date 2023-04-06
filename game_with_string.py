word = input()
stack = []
r, i = 0,0

for c in word:
    if len(stack) > 0:
        if stack[-1] == c:
            r += 1
            stack.pop()
        else:
            stack.append(c)
    else:
        stack.append(c)

if r % 2 == 0:
    print("No")
else:
    print("Yes")