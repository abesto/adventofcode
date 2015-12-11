a = ord('a')
z = ord('z')
forbidden = [ord(c) for c in 'iol']


def step_char(n):
    global a, z
    if n == z:
        return a, True
    return n + 1, False


def step_pw(ns, pos=None):
    if pos is None:
        pos = len(ns) - 1
    overflow = True
    while overflow:
        old = ns[pos]
        new, overflow = step_char(old)
        ns[pos] = new
        pos -= 1
    return ns


def is_valid(ns):
    has_straight = False
    found_pair = 0  # 0: non yet. 1: found one, on previous character. 2: found one, NOT on previous letter. 3: found two.
    memory = []
    for n in ns:
        if n in forbidden:
            return False, 'Forbidden %s' % chr(n)
        if len(memory) > 0 and n == memory[-1]:
            found_pair += 1
        elif found_pair == 1:
            found_pair += 1
        if len(memory) > 1 and memory[-2] + 1 == memory[-1] and memory[-1] + 1 == n:
            has_straight = True
        memory.append(n)
        while len(memory) > 3:
            memory.pop(0)
    if found_pair != 3:
        return False, 'found_pair = %s' % found_pair
    if not has_straight:
        return False, 'no straight'
    return True, ''


def read_pw(s):
    return [ord(c) for c in s]

def write_pw(ns):
    return ''.join([chr(n) for n in ns])

def next_valid(pw):
    ns = step_pw(read_pw(pw))
    while not is_valid(ns)[0]:
        # Optimize out lost of tests for forbidden chars
        for n in xrange(len(ns)):
            while ns[n] in forbidden:
                ns = step_pw(ns, n)
        ns = step_pw(ns)
    return write_pw(ns)


if __name__ == '__main__':
    #print 'xz ->', write_pw(step_pw(read_pw('xz')))
    print 'hijklmmn', is_valid(read_pw('hijklmmn'))
    print 'abbceffg', is_valid(read_pw('abbceffg'))
    print 'abbcegjk', is_valid(read_pw('abbcegjk'))

    for pw in ['abcdefgh', 'ghijklmn', 'hepxcrrq']:
        print '%s ->' % pw, next_valid(pw)
