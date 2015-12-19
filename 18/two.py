import time
import sys


class Grid(object):
    def __init__(self, rows, columns):
        self.data = []
        self.rows = rows
        self.columns = columns
        for n in xrange(rows):
            self.data.append([0] * columns)

    def get(self, row, column):
        if row in (0, self.rows - 1) and column in (0, self.columns - 1):
            return 1
        if row < 0 or row >= self.rows or column < 0 or column >= self.columns:
            return 0
        return self.data[row][column]

    def set(self, row, column, value):
        self.data[row][column] = value

    def neighbors_on(self, row, column):
        return sum(self.get(r, c)
                   for r in [row - 1, row, row + 1]
                   for c in [column - 1, column, column + 1]
                   if not (row == r and column == c))

    def mutate(self):
        new = Grid(self.rows, self.columns)
        for row in xrange(self.rows):
            for column in xrange(self.columns):
                if self.get(row, column):
                    new.set(row, column, self.neighbors_on(row, column) in (2, 3))
                else:
                    new.set(row, column, self.neighbors_on(row, column) == 3)
        return new

    def count_on(self):
        return sum(self.get(row, column)
                   for row in xrange(self.rows)
                   for column in xrange(self.columns))

    def __repr__(self):
        reprs = '.#'
        return '\n'.join(
            ''.join(reprs[self.get(r, c)] for c in xrange(self.columns))
            for r in xrange(self.rows)
        )


def read_input(filename):
    with open(filename, 'r') as f:
        lines = [line.strip() for line in f]
    assert all(len(line) == len(lines[0]) for line in lines)
    grid = Grid(len(lines), len(lines[0]))
    for row, line in enumerate(lines):
        for column, char in enumerate(line):
            grid.set(row, column, {'.': 0, '#': 1}[char])
    return grid


def iterate(grid, n):
    steps = [grid]
    for iteration in xrange(n):
        steps.append(steps[-1].mutate())

    print 'Lights on after %d steps: %d' % (n, steps[-1].count_on())
    #for i, step in enumerate(steps):
        #print 'Step %d' % i
        #print step
        #print
        #time.sleep(1)
        #sys.stdout.write('\r' + ('\033[1A' * (len(steps[0].data) + 1)))
    #print steps[-1]


iterate(read_input('input.example'), 5)
iterate(read_input('input'), 100)
