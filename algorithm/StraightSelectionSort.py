# Thuat toan sap xep lua chon truc tiep
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
    min = i
    for j in range(i + 1, len(a)):
        if int(a[min]) > int(a[j]):
            print("Set Min({}[{}]) = {}[{}]".format(min, a[min], j, a[j]))
            min = j
        else:
            print("Keep Min({}[{}])".format(min, a[min]))
    if min != i:
        swapValue(a, i, min)
    print("Arr:", a)

print("Result arr: ", a)