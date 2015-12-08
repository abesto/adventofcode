class Circuit(object):
    def __init__(self):
        self.nodes = {}

    def set(self, name, node):
        if name in self.nodes:
            assert isinstance(self.nodes[name], NodePromise)
            self.nodes[name].set(node)
        else:
            self.nodes[name] = node

    def get(self, name):
        try:
            return Const(int(name))
        except ValueError:
            if name not in self.nodes:
                self.nodes[name] = NodePromise(name)
            return self.nodes[name]


class Node(object):
    op_functions = {
        'AND': lambda x, y: x & y,
        'OR': lambda x, y: x | y,
        'NOT': lambda x: ~ x,
        'LSHIFT': lambda x, y: x << y,
        'RSHIFT': lambda x, y: x >> y
    }

    def __init__(self, name, function):
        self.inputs = []
        self.function = function
        self.name = name
        self.value = None

    def connect(self, input_node):
        self.inputs.append(input_node)

    def arity(self):
        raise NotImplemented()

    def evaluate(self):
        if self.value is None:
            args = [i.evaluate() for i in self.inputs]
            self.value = self.function(*args)
        print '%s (%s)' % (self, self.value)
        return self.value

    def reset(self):
        self.value = None

    def __str__(self):
        return self.name

    @classmethod
    def from_opname(cls, line, opname):
        return cls(line, cls.op_functions[opname])


class NodePromise(object):
    def __init__(self, name):
        self.value = None
        self.name = name

    def set(self, value):
        self.value = value

    def evaluate(self):
        return self.value.evaluate()

    def reset(self):
        self.value.reset()

    def __str__(self):
        if self.value == None:
            return 'Promise(%s)' % self.name
        return str(self.value)


class Const(Node):
    def __init__(self, const):
        super(Const, self).__init__(str(const), lambda: const)


def parse_line(line, circuit):
    words = line.split(' ')
    output = words.pop()
    assert words.pop() == '->'
    if len(words) == 1:
        return circuit.get(words[0]), output
    if len(words) == 2:
        assert words[0] == 'NOT'
        node = Node.from_opname(line, words.pop(0))
    else:
        node = Node.from_opname(line, words.pop(1))
    for arg in words:
        node.connect(circuit.get(arg))
    return node, output


def read_input():
    circuit = Circuit()
    with open('input', 'r') as f:
        for line in f:
            node, output = parse_line(line.strip(), circuit)
            circuit.set(output, node)
    return circuit
