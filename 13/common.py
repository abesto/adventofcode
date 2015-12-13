from collections import defaultdict
from itertools import permutations


def read_input():
    prefs = defaultdict(lambda: defaultdict(int))
    with open('input', 'r') as f:
        for line in f:
            words = line.strip('\n.').split(' ')
            prefs[words[0]][words[10]] = int(words[3]) * {'gain': 1, 'lose': -1}[words[2]]
    return prefs


def solve(prefs):
    best = 0
    for setup in permutations(prefs.iterkeys()):
        current = 0
        for n in xrange(-1, len(setup) - 1):
            current += prefs[setup[n]][setup[n + 1]]
            current += prefs[setup[n + 1]][setup[n]]
        if current > best:
            best = current
    return best
