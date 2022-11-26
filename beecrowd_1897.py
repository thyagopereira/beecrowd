import pdb

def op0(n):
  return n * 2

def op1(n):
  return n * 3

def op2(n):
  return n/ 2

def op3(n):
  return n /3

def op4(n):
  return n + 7

def op5(n):
  return n - 7

def apply(comb ,n):
  applied = n

  for i in comb:
    if i != None:
      if i == 0:
        applied = op0(applied)
      elif i == 1:
        applied = op1(applied)
      elif i == 2:
        applied = op2(applied)
      elif i == 3:
        applied = op3(applied)
      elif i == 4:
        applied = op4(applied)
      elif i == 5:
        applied = op5(applied)

  return applied


result = []

def perm(comb, n, m, i):

  a = apply(comb, n[0]) # applica a combinação para poder comparar  

  if i == 6 :
    return
  
  if m == a : # Verifica se essa solução é possivel
    result.append(comb) 
  else: 
    for j in range(0,6):
      comb[i] = j
      perm(comb, n, m, i + 1)


def main():

  n, m = [int(x) for x in input().split(" ")]
  n = [n]

  comb = [None] * 6   
  i = 0
  perm(comb, n, m, i)

  print(result)
  
main()






