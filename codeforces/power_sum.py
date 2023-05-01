# Questão simples de backtracking

def power_sum(x, n, num):
    val = x - pow(num, n)

    if val < 0:
        return 0
    elif val == 0:
        return 1
    else:
        return power_sum(val, n, num+1) + power_sum(x, n, num + 1)
    


x = int(input())
n = int(input())

print(power_sum(x, n, 1))