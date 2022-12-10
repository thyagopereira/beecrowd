W = 50

def toys(j):

  pack = []
  while j > 0:

    t, w = [int(x) for x in input().split(" ")]
    pack.append({"v": t, "w": w})

    j -= 1
  
  return pack

def main():
  n = int(input()) 

  while n > 0:
    j = int(input())

    result = []
    packs = toys(j)
    a = mochila(packs, result, 0, W)

    print(result)
    print(a)
    n -= 1

main()