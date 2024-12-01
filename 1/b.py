import sys

lefts = []
rights = []

for line in sys.stdin:
    bits = line.strip().split('   ')
    lefts.append(int(bits[0]))
    rights.append(int(bits[1]))

counts = {}
for r in rights:
    counts[r] = counts.get(r, 0) + 1

print(sum(l * counts.get(l, 0) for l in lefts))