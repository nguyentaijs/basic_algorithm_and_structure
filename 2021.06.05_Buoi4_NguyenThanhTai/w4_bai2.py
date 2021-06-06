_inputStr = input("Input separated by space char: ")

_input = _inputStr.split()
_output = []
print("Input : ", _input)

def swapValue(_input, i, j):
	temp = _input[i]
	_input[i] = _input[j]
	_input[j] = temp
	print("Swap {}[{}] with {}[{}]".format(i, _input[i], j, _input[j]))

for i in range(len(_input) - 1):
	print('---[i = {}]---'.format(i))
	for j in range(i+1, len(_input)):
		if int(_input[j]) > int(_input[i]):
			swapValue(_input, j, i)
	print("Keep {}[{}] & {}[{}]".format(i, _input[i], j, _input[j]))
	print("New input:", _input)

print("Output: ", _input)