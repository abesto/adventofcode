import itertools
import json


def value_iterator(o):
    if isinstance(o, dict):
        return itertools.chain(*[value_iterator(item) for item in o.itervalues()])
    if isinstance(o, list):
        return itertools.chain(*[value_iterator(item) for item in o])
    return [o]


def reduce_values(o):
    val = 0
    for item in value_iterator(o):
        if isinstance(item, int):
            val += item
    return val


if __name__ == '__main__':
    with open('input', 'r') as f:
        data = json.load(f)
    print reduce_values(data)
