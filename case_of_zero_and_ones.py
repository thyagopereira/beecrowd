n = int(input())

string = input()
aux = {1: 0 , 0: 0}

for s in string:
    if s == '0':
        aux[0] += 1
    else:
        aux[1] += 1

if aux[0] - aux[1] != 0:
    print(abs(aux[0] - aux[1]))
else:
    print(0) 
