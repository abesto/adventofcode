result = 0
with open('input', 'r') as f:
    for line in f:
        line = line.strip()
        result += len(line) - len(eval(line))
print result
