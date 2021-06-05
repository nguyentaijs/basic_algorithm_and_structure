result = ''

for _num in range(1000, 3000 + 1):
	if _num % 2 == 0:
		result += str(_num)
		result += ', '

print(result)