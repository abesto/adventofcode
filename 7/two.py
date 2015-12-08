from common import read_input, Const

circuit = read_input()
a_val = circuit.get('a').evaluate()
for node_name in circuit.nodes.keys():
    circuit.get(node_name).reset()
circuit.get('b').value = Const(a_val)
print circuit.get('a').evaluate()
