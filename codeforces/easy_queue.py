t = int(input())
queue = []
while t > 0:
    
    q = [int(x) for x in input().split(" ")]

    if q[0] == 1:
        queue.append(q[1])
    elif q[0] == 2:
        if len(queue) > 0 :
            queue.pop(0)
    elif q[0] == 3:
        if len(queue) > 0:
            print(queue[0])
        else:
            print("Empty!")
    t -= 1