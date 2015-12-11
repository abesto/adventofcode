ops = {
    'on': lambda x: x + 1,
    'off': lambda x: max(x - 1, 0),
    'toggle': lambda x: x + 2
}


def parse_line(line):
    words = line.strip().split(' ')
    op = words.pop(0)
    if op != 'toggle':
        op = words.pop(0)
    from_row, from_column = map(int, words.pop(0).split(','))
    words.pop(0)
    to_row, to_column = map(int, words.pop(0).split(','))
    assert len(words) == 0
    return ops[op], from_row, from_column, to_row, to_column


field = [[False] * 1000 for n in xrange(10000)]

with open('input', 'r') as f:
    for line in f:
        op, from_row, from_column, to_row, to_column = parse_line(line)
        for row in xrange(from_row, to_row + 1):
            for column in xrange(from_column, to_column + 1):
                field[row][column] = op(field[row][column])

result = 0
for row in field:
    result += sum(row)

print result
