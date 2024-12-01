import sys

lefts = []
rights = []

for line in sys.stdin:
    bits = line.strip().split('   ')
    lefts.append(int(bits[0]))
    rights.append(int(bits[1]))

lefts.sort()
rights.sort()

print(sum(abs(x[0]-x[1]) for x in zip(lefts, rights)))
