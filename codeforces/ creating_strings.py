# C - Creating Strings

s = input()
resp = {}
def generate(s, l, r):
    if l == r :
        global resp
        string = ''.join(s)
        if string not in resp.keys():
            resp[string] = ''
    else:
        for i in range(l, r):
            s[l], s[i] = s[i], s[l]
            generate(s, l+1, r)
            s[l], s[i] = s[i], s[l]

generate(list(s), 0, len(s))
palavras = resp.keys()

print(len(palavras))
for word in sorted(palavras):
    print(word)