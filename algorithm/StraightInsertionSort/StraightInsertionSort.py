# Thuat toan sap xep chen truc tiep
arrayStr = input("Input separated by space char: ")

a = arrayStr.split()
print("Arr : ", a)


def swapValue(a, i, j):
  print("Swap {}[{}] with {}[{}]".format(i, a[i], j, a[j]))
  temp = a[i]
  a[i] = a[j]
  a[j] = temp


for i in range(1, len(a)):
  print('---[i = {}]---'.format(i))
  position = i
  currentVal = a[i]
  while position > 0 and int(a[position - 1]) > int(currentVal):
    a[position] = a[position - 1]
    position = position - 1
  a[position] = currentVal
  print("Arr:", a)

print("Result arr: ", a)