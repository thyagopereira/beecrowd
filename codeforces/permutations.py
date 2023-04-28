n = int(input())
par, im = [], []

for i in range(1, n + 1):
    if i % 2 == 0:
        par.append(i)
    else:
        im.append(i)
if n > 1:
	if par[len(par) -1] - im[0] != abs(1):
		for i in par + im:
			print(i, end=" ")         
	else:
		print("NO SOLUTION")

else:
    print(n)
    
