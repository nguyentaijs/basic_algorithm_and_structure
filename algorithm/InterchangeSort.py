# Thuat toan sap xep doi cho truc tiep
arrayStr = input("Input separated by space char: ")

a = arrayStr.split()
print("Arr : ", a)


def swapValue(a, i, j):
    print("Swap {}[{}] with {}[{}]".format(i, a[i], j, a[j]))
    temp = a[i]
    a[i] = a[j]
    a[j] = temp


for i in range(len(a) - 1):
    print('---[i = {}]---'.format(i))
    for j in range(i + 1, len(a)):
        if int(a[i]) > int(a[j]):
            swapValue(a, i, j)
        else:
            print("Keep {}[{}] & {}[{}]".format(i, a[i], j, a[j]))
    print("Arr:", a)

print("Result arr: ", a)