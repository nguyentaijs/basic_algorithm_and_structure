_inputStr = input("Input separated by space char: ")

_input = _inputStr.split()
_output = []
print("Input : ", _input)

def swapValue(_input, i, max):
	temp = _input[i]
	_input[i] = _input[max]
	_input[max] = temp

for i in range(len(_input) - 1):
	print('---[i = {}]---'.format(i))
	max = i
	for j in range(i+1, len(_input)):
		if int(_input[j]) >= int(_input[max]):
			max = j
		print("j = ", j)
		print("max = ", max)
		print("New input:", _input)
	swapValue(_input, i, max)


print("Output: ", _input)