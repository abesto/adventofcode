import re
import cProfile
from copy import copy


def all_indices(s, substring):
    offset = 0
    while True:
        index = s.find(substring, offset)
        if index == -1:
            break
        yield index
        offset = index + 1


inf = float('inf')
seen = set()


def contract(molecule, replacements):
    global inf, seen
    #print '>', molecule
    if molecule == 'e':
        return 0, molecule
    if 'e' in molecule or molecule in seen:
        return inf, molecule
    retval = inf, molecule
    seen.add(molecule)

    for left, cost, right in replacements:
        for start in all_indices(molecule, left):
            pre = molecule[:start]
            post = molecule[start + len(left):]
            new = pre + right + post

            #print left, cost, right, new
            rest_cost, rest_result = contract(new, replacements)
            retval = min(retval, (cost + rest_cost, rest_result), key=lambda p: p[0])
            if retval[0] < inf:
                return retval

    #print '<', molecule, retval
    print '%s\t%s\t%s' % (len(seen), len(molecule), molecule)
    return retval


def read_input(filename):
    replacements = []  # [(from, to, cost)]
    with open(filename, 'r') as f:
        lines = [line.strip() for line in f]
    input_molecule = lines.pop()
    assert lines.pop() == ''
    for line in lines:
        item = line.split(' => ')
        replacements.append((item[1], 1, item[0]))
    return input_molecule, replacements


def expand(molecule, replacements):
    molecules = set()
    for replacement in replacements:
        #print replacement
        for match in re.finditer(replacement[0], molecule):
            pre = molecule[:match.start()]
            post = molecule[match.end():]
            new = pre + replacement[2] + post
            #print molecule, pre, post, replacement, new
            molecules.add(new)
    return molecules


def expand_replacements(replacements):
    for (long, cost, short) in copy(replacements):
        for longer in expand(long, replacements):
            new = (longer, cost + 1, short)
            if new not in replacements:
                replacements.append((longer, cost + 1, short))

#print contract(*read_input('input.example3'))
#print contract(*read_input('input.example4'))
#cProfile.run("contract(*read_input('input'))")

molecule, replacements = read_input('input')
#expand_replacements(replacements)
#expand_replacements(replacements)

print contract(molecule, replacements)[0] + 1
