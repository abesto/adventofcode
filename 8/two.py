result = 0
with open('input', 'r') as f:
    for line in f:
        memory = line.strip()
        code = memory.encode('string_escape').replace('"', '""') + '""'  # don't actually need to escape, just get the char count right
        result += len(code) - len(memory)
print result
