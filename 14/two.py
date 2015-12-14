from common import read_input


def race(reindeer, seconds):
    n = 0
    while n < seconds:
        # Move
        for r in reindeer:
            r.tick()
        n += 1
        # Find leaders, give them a point
        max_position = 0
        leaders = []
        for r in reindeer:
            if r.position > max_position:
                max_position = r.position
                leaders = []
            if r.position == max_position:
                leaders.append(r)
        for r in leaders:
            r.score += 1
    return reindeer


print 'EXAMPLE'
result = race(read_input('input.example'), 1000)
print result
print 'Won:', max(result, key=lambda r: r.score)

print 'REAL'
result = race(read_input('input'), 2503)
print result
print 'Won:', max(result, key=lambda r: r.score)
