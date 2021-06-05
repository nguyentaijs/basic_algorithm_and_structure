input = input("Please input: ")
lowerCount = 0
upperCount = 0
numberCount = 0

for inputChar in input:
	if inputChar.isdigit():
		numberCount += 1
	elif inputChar.islower():
		lowerCount += 1
	elif inputChar.isupper():
		upperCount += 1
	else:
		print("Unrecognized type {}", inputChar)

print("Total numbers: ", numberCount)
print("Total lowers: ", lowerCount)
print("Total uppers: ", upperCount)