import re

# 1. Get user input
passwd = input("Enter password: ")
print("Your password is: ", passwd)

# 2. Check password is valid
# 2.1. Regex with python
pattern = re.compile("^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[\$\#\@]).*$")

result = pattern.search(passwd)

# 2.2. Check length
if result:
	result = len(passwd) <= 12
	if result:
		result = len(passwd) >= 6

# 3. Print result
if result:
	print("Matched")
else:
	print("Not match")
