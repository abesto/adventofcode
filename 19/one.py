import re


def expand(molecule, replacements):
    molecules = set()
    for replacement in replacements:
        #print replacement
        for match in re.finditer(replacement[0], molecule):
            pre = molecule[:match.start()]
            post = molecule[match.end():]
            new = pre + replacement[1] + post
            #print molecule, pre, post, replacement, new
            molecules.add(new)
    return molecules


def read_input(filename):
    replacements = []
    with open(filename, 'r') as f:
        lines = [line.strip() for line in f]
    input_molecule = lines.pop()
    assert lines.pop() == ''
    for line in lines:
        replacements.append(line.split(' => '))
    return input_molecule, replacements


if __name__ == '__main__':
    print len(expand(*read_input('input.example')))
    print len(expand(*read_input('input.example2')))
    print len(expand(*read_input('input')))
