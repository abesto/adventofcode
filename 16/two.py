from operator import eq, gt, lt
from collections import defaultdict


def read_tape(filename):
    data = {}
    with open(filename, 'r') as f:
        for line in f:
            key, value = line.strip().split(': ')
            assert key not in data
            data[key] = int(value)
    return data


def read_input(filename):
    aunts = []
    with open(filename, 'r') as f:
        for line in f:
            words = line.strip().replace(',', '').replace(':', '').split(' ')
            assert words.pop(0) == 'Sue'
            assert int(words.pop(0)) == len(aunts) + 1
            aunt = {}
            while len(words) > 0:
                key = words.pop(0)
                value = int(words.pop(0))
                aunt[key] = value
            aunts.append(aunt)
    return aunts


operators = defaultdict(lambda: eq)
operators['cats'] = operators['trees'] = lt
operators['pomeranians'] = operators['goldfish'] = gt

tape = read_tape('tape')
aunts = read_input('input')
for aunt_index, aunt in enumerate(aunts):
    is_the_right_aunt = all(
        key not in aunt or operators[key](tape[key], aunt[key])
        for key in tape.iterkeys()
    )
    if is_the_right_aunt:
        print tape, aunt, aunt_index + 1
