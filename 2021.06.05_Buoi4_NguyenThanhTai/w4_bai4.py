balance = int(input("Init balance: "))
transactionStr = input("Input transaction separated by space char: ")

transactions = transactionStr.split()
isWithDraw = False

for item in transactions:
	if item.isalpha():
		if item.lower == "w":
			isWithDraw = True
		else:
			isWithDraw = False
	if item.isdigit():
		if isWithDraw:
			if balance < int(item):
				print("Can not withdraw, Balance {} < Withdraw {}, skip this transaction".format(balance, item))
				continue
			balance -= int(item)
			print("Withdraw {}. Balance = {}".format(item, balance))
		else:
			balance += int(item)
			print("Deposit {}. Balance = {}".format(item, balance))

