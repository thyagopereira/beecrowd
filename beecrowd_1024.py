
def move_3(s):
  result = ""

  for char in s:
    if char.isalpha():
      result += chr(ord(char) + 3)
    else:
      result += char
  
  return result

def reverse(s):
  result = ""
  for i in range(len(s) - 1, -1, -1):
    result += s[i]

  return result

def move_back(s):
  result = ""
  middle = int(len(s) / 2)

  for i in range(0, len(s)):
    if i >= middle:
      result += chr(ord(s[i]) -1)
    else:
      result += s[i]

  return result

def main():
  t = int (input())

  while(t > 0):
    
    text = input()
    a = move_3(text)
    b = reverse(a)
    print(move_back(b))

    t -= 1
    
main()