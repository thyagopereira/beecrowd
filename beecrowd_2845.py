def mdc(a,b):
  while (b != 0):
    resto = a % b
    a = b
    b = resto
  
  return a

def coprimo(a,b):
  return mdc(a,b) == 1

def party(letter, result, n, hp):
  valid = True

  for number in letter:
    if not coprimo(number, hp):
      valid = False
      party(letter, result, n, hp + 1)

  if valid:
    result.append(hp)

def main():
  n = int(input())

  result = []
  letter = [int (x) for x in input().split(" ")]

  hp =  max(letter) + 1
  party(letter, result, n, hp)

  print(result[0])

main()