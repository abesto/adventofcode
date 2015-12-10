class Vertex(object):
    def __init__(self, name):
        self.edges = []
        self.name = name


class Edge(object):
    def __init__(self, a, b, distance):
        self.a = a
        self.b = b
        self.distance = distance

    def get_neighbor(self, vertex):
        assert self.touches(vertex)
        if vertex == self.a:
            return self.b
        return self.a

    def touches(self, vertex):
        return vertex is self.a or vertex is self.b


class Graph(object):
    def __init__(self):
        self.vertices = {}

    def get_vertex(self, name):
        if name not in self.vertices:
            self.vertices[name] = Vertex(name)
        return self.vertices[name]

    def connect(self, name_a, name_b, distance):
        a = self.get_vertex(name_a)
        b = self.get_vertex(name_b)
        edge = Edge(a, b, distance)
        a.edges.append(edge)
        b.edges.append(edge)


# build graph
g = Graph()
with open('input', 'r') as f:
    for line in f:
        words = line.strip().split(' ')
        distance = int(words.pop())
        assert words.pop() == '='
        name_b = words.pop()
        assert words.pop() == 'to'
        name_a = words.pop()
        g.connect(name_a, name_b, distance)


def solve(operator):
    # brute-force
    def visit(current, visited):
        try:
            return operator(edge.distance + visit(edge.get_neighbor(current), visited.union(set([current])))
                            for edge in current.edges
                            if edge.touches(current)
                            and edge.get_neighbor(current) not in visited
            )
        except ValueError as e:
            return 0


    print operator(visit(vertex, set()) for vertex in g.vertices.itervalues())
