#Loop Musical

def get_picos(wave, result):

  for i in range(len(wave)):
    if i == 0 :
      now, prev, nex = wave[i], wave[-1], wave[i +1]
    elif i == (len(wave) -1 ) :
      now, prev, nex =  wave[i], wave[i -1], wave[0]
    else:
      now, prev, nex = wave[i], wave[i -1], wave[i + 1]

    delta_i = now - prev 
    delta_j = nex - now

    if not ( ( delta_i * delta_j ) > 0 ):
      result.append(1)

def main():

  while (True):

    n = int(input())

    if n == 0:
      break

    wave = [int (x) for x in input().split(" ")]

    result = []
    get_picos(wave, result)

    print(len(result))
      
main()