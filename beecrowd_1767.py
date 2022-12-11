w = 50

def toys(j):

  pack = []
  pack.append(0)
  while j > 0:

    t, w = [int(x) for x in input().split(" ")]
    pack.append({"v": t, "w": w})

    j -= 1
  
  return pack

def create_table(line, collum):
  table = []
  for i in range(line):
    table.append([])

  for j in range(line):
    for i in range(collum):
      table[j].append([])
  
  return table

def table(packs):
  table  = create_table(len(packs), w)

  for i in range(len(packs)):
    for j in range(w):
      if i == 0 or j == 0:
        table[i][j] = 0
      elif j <= packs[i]["w"]:
        table[i][j] = max(packs[i]["v"], table[i - 1][j])
      else:
        left_cap = j - packs[i]["w"]
        val = table[i - 1][left_cap]
        # ver se cabe antes de botar na pica do maximo ---- teria que ser ordenado por Peso ?
        
        table[i][j] = max(packs[i]["v"] + val, table[i -1][j])

  return table

def read_table(table, packs):
  brq = table[len(table) -1][w -1]
  inserted = []

  c = w - 1
  l = len(packs) -1

  while True:
    if l == 1:
      break
    else:
      if table[l][c] != table[l - 1][c]:
        inserted.append(packs[l])
        new_c = c - packs[l]["w"]
        l = l - 1
        c = new_c
      
  return inserted, brq

def total_peso(inserted):
  total = 0
  for elem in inserted:
    total += elem["w"]

  return total

def print_result(inserted, brq, packs):
  total = total_peso(inserted)
  left  = len(packs) - len(inserted)

  print( str(brq) + " brinquedos")
  print( "Peso: " + str(total) + " kg")
  print("sobra(m) " + str(left) + " pacote(s)")

def main():
  n = int(input()) 

  while n > 0:
    j = int(input())

    packs = toys(j)
    result = table(packs)
    print(result)
    inserted, brq = read_table(result, packs)

    print_result(inserted, brq, packs)
    n -= 1

main()