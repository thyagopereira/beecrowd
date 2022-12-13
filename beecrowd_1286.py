def entry(n, val, peso):
	for _ in range(n):
		v, p = [int (x) for x in input().split(" ")]
		val.append(v)
		peso.append(p)


def z(itens, cap, val, peso, table):

	#Condição de contorno
	if (itens == 0) or (cap == 0):
		if str(itens) not in table.keys():
			table[str(itens)] = {}
		table[str(itens)][str(cap)] = 0
		return 0

	#Se não caiu nas condições de contorno, verificaremos se na tabela já existe um valor associado ao calculo
	elif (str(itens) in table.keys()):
		if str(cap) in table[str(itens)].keys():
			#Existindo valor associado ao cáculo recuperamos ele. (Nesse caso apenas o val max é importante)
			return table[str(itens)][str(cap)]
	else:
		# Não existindo esse valor, ele precisa ser calculado, Aqui entra a definição da função recursiva que é o primeiro passo para a Programação dinâmica
		for i in range(1,itens + 1):
			if i not in table.keys():
				table[str(i)] = {}
			for j in range(cap + 1):
				# Se cabe na mochila, então vamos considerar com e sem o valor;
				if peso[i] <= j:
					table[str(i)][str(j)] = max(z(i -1 , j, val, peso, table), val[i] + z(i -1 , j - peso[i], val, peso, table))
				else:
					# Uma vez o item I não cabe na mochila. consideramos sem ele
					table[str(i)][str(j)] = z(i -1, j, val, peso, table)

		return table[str(itens)][str(cap)]


def main():

	while True:
		n =  int(input())

		if n == 0:
			break

		val = [] 
		peso = []
		cap = int(input())

		val.append(0) # Ajuste para começar do Item 1, começar no Item 0, faz apelar para -1 elementos
		peso.append(0) # Ajuste para começar do Item 1, começar no Item 0, faz apelar para -1 elementos

		entry(n, val, peso)

		table = {} # [ qtd de itens considerados, capacidade considerada]
		
		finish = z(len(val) -1, cap, val, peso, table)
		print("{} min.".format(finish))
		
main()