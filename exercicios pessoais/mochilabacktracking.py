pesos , values, my_bag = [15, 5, 10, 5], [20, 30, 50, 10], [None] * 4
max_w = 16

def calcVal(values, bag):
    total = 0 

    for i in range(len(bag)):
        if bag[i] == 1:
            total += values[i]

    return total

def calcW(wei, bag):
    total = 0 

    for i in range(len(bag)):
        if bag[i] == 1:
            total += wei[i]

    return total

best_val = 0 
def mochila(weights, values, my_bag, max_w, i):
    global best_val
    if i == len(my_bag) : 
        # Decidi se adiciono ou não todos os itens. O caso base é chegar numa solucao valida
        # Dada uma soluçao valida, é a mulhor possivel ? 
        
        sol_val = calcVal(values, my_bag) 
        if  sol_val > best_val:
            best_val = sol_val
            print(my_bag, best_val)    
    else:
        w = calcW(weights, my_bag)
        if w + weights[i] < max_w:
            my_bag[i] = 1
            mochila(weights, values, my_bag, max_w, i + 1)
        
        my_bag[i] = 0
        mochila(weights, values, my_bag, max_w, i + 1)

mochila(pesos, values, my_bag, max_w, 0)