def step(s):
    out = ''
    offset = 0
    l = len(s)
    while offset < l:
        c = s[offset]
        n = 1
        while offset + n < l and  s[offset + n] == c:
            n += 1
        out += '%s%s' % (n, c)
        offset += n
    return out


def solve(s, steps):
    n = 0
    while n < steps:
        s = step(s)
        n += 1
    return s

