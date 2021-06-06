_inputStr = input("Input separated by space char: ")

_input = _inputStr.split()
row = int(_input[0])
col = int(_input[1])
_array = []

for i in range(0, row):
	_array.append([])
	for j in range(0, col):
		_array[i].append(i * j)

print("Output: ", _array)