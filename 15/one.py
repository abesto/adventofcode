# Brute force, because if it fits, I sits.
import operator

class Ingredient(object):
    def __init__(self, name, axes):
        self.name = name
        self.axes = axes

    def times(self, n):
        return [n * m for m in self.axes]


class Recipe(object):
    def __init__(self):
        self.scores = [0] * 4

    def add(self, n, ingredient):
        scores = ingredient.times(n)
        for n in xrange(len(self.scores)):
            self.scores[n] += scores[n]

    def score(self):
        return reduce(operator.mul, [max(0, s) for s in self.scores], 1)


def read_input(filename):
    ingredients = []
    with open(filename, 'r') as f:
        for line in f:
            words = line.strip().split(' ')
            ingredients.append(Ingredient(
                words[0].strip(':'),
                [int(words[n].strip(',')) for n in [2, 4, 6, 8]]
            ))
    return ingredients


def brute_force(ingredients):
    best = -1
    if len(ingredients) == 4:
        options = [
            [a, b, c, d]
            for a in xrange(101)
            for b in xrange(101 - a)
            for c in xrange(101 - a - b)
            for d in xrange(101 - a - b - c)
            if a + b + c + d == 100
        ]
    elif len(ingredients) == 2:
        options = [
            [a, b]
            for a in xrange(101)
            for b in xrange(101 - a)
            if a + b == 100
        ]
    else:
        raise Exception("I'm not prepared for having this many (or this few) ingredients in my kitchen. I think I'm just going to give up.")
    for option in options:
        recipe = Recipe()
        for ingredient_index, n in enumerate(option):
            recipe.add(n, ingredients[ingredient_index])
        score = recipe.score()
        if score > best:
            print option, score
            best = score
    return best


print 'EXAMPLE ONE'
print brute_force(read_input('input.example'))
print

print 'THE REAL DEAL'
print brute_force(read_input('input'))

