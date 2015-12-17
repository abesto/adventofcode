from itertools import chain
from copy import copy


def ways(unused_containers, used_containers, eggnog):
    if eggnog == 0:
        return [used_containers]
    if eggnog < 0:
        return []
    return chain(*(ways(unused_containers[i+1:], used_containers + [container], eggnog - container) for i, container in enumerate(unused_containers)))


def solve(containers, eggnog):
    myways = ways(containers, [], eggnog)
    minlen = float('inf')
    solutions = []
    for way in myways:
        if len(way) == minlen:
            solutions.append(way)
        elif len(way) < minlen:
            solutions = [way]
            minlen = len(way)
    return len(solutions)

print 'EXAMPLE'
print solve([20, 15, 10, 5, 5], 25)

print 'REAL'
with open('input', 'r') as f:
    containers = [int(line.strip()) for line in f]
print solve(containers, 150)
