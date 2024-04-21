def runLengthEncoding(string):
	encodedStringCaracthers = [] 
	currentRunLength = 1

	for i in range(1, len(string)):
		curr, prev = string[i], string[i -1]

		if curr != prev or currentRunLength == 9 :
			encodedStringCaracthers.append(str(currentRunLength))
			encodedStringCaracthers.append(prev)

			currentRunLength = 0 
		
		currentRunLength += 1

	encodedStringCaracthers.append(str(currentRunLength))
	encodedStringCaracthers.append(string[len(string) -1])

	return "".join(encodedStringCaracthers)