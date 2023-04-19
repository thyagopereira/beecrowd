n = int(input())

mailbox = [int (x) for x in input().split(" ")]
count, ones = 0,0

for m in mailbox:
	if m == 1:
		ones += 1

i = 0
while ones > 0 and i < len(mailbox):
	
	if mailbox[i] == 1 :
		ones -= 1 
		if ones > 0:
			count += 2
			if mailbox[i + 1] == 1 :
				count -= 1 
		else:
			count += 1
	i += 1

print(count)	
