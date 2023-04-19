n = int(input())

register = {}
while n > 0:
    name = input()
    if name not in register.keys():
        register[name] = 1
        print("OK")
    else:
        novo = name + str(register[name])
        register[name] += 1
        print(novo)

    n -= 1
