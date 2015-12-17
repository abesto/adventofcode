from copy import copy


def ways(containers, eggnog):
    if eggnog == 0:
        return 1
    if eggnog < 0:
        return 0
    return sum(ways(containers[i+1:], eggnog - container) for i, container in enumerate(containers))


print 'EXAMPLE'
print ways([20, 15, 10, 5, 5], 25)

with open('input', 'r') as f:
    containers = [int(line.strip()) for line in f]

print 'REAL'
print ways(containers, 150)
