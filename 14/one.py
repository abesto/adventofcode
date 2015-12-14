from common import read_input, leader


def race(reindeer, seconds):
    n = 0
    while n < seconds:
        for r in reindeer:
            r.tick()
        n += 1
    return reindeer


def leader(reindeer):
    return max(reindeer, key=lambda r: r.position)


print 'EXAMPLE'
result = race(read_input('input.example'), 1000)
print result
print 'Won: %s' % leader(result)

print 'REAL'
result = race(read_input('input'), 2503)
print result
print 'Won: %s' % leader(result)
