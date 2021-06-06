_inputStr = input("Input separated by space char: ")

_input = _inputStr.split()
print("Input : ", _input)

def swapValue(_input, i, j):
	temp = _input[i]
	print("Swap {}[{}] with {}[{}]".format(i, _input[i], j, _input[j]))
	_input[i] = _input[j]
	_input[j] = temp

for i in range(len(_input) - 1):
	print('---[i = {}]---'.format(i))
	swapped = False
	for j in range(0, len(_input) - i - 1):
		print('[j = {}]'.format(j))
		if int(_input[j]) < int(_input[j + 1]):
			swapValue(_input, j, j + 1)
			swapped = True
		else:
			print("Keep {}[{}] & {}[{}]".format(j, _input[j], j + 1, _input[j + 1]))
	if swapped == False:
		print("Array is already sorted")
		break
	print("New input:", _input)

print("Output: ", _input)