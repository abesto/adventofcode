# FSM FTW
class State(object):
    def __init__(self, reindeer):
        self.reindeer = reindeer

    def tick(self):
        raise NotImplementedError()


class StateFly(State):
    def __init__(self, reindeer):
        super(StateFly, self).__init__(reindeer)
        self.reindeer.ticks = self.reindeer.flight_time

    def tick(self):
        self.reindeer.position += self.reindeer.speed
        self.reindeer.ticks -= 1
        if self.reindeer.ticks == 0:
            return StateRest(self.reindeer)
        return self


class StateRest(State):
    def __init__(self, reindeer):
        super(StateRest, self).__init__(reindeer)
        self.reindeer.ticks = self.reindeer.rest_time

    def tick(self):
        self.reindeer.ticks -= 1
        if self.reindeer.ticks == 0:
            return StateFly(self.reindeer)
        return self


class Reindeer(object):
    def __init__(self, name, speed, flight_time, rest_time):
        self.name = name
        self.speed = speed
        self.flight_time = flight_time
        self.rest_time = rest_time
        self.position = 0
        self.state = StateFly(self)
        self.score = 0

    def tick(self):
        self.state = self.state.tick()

    def __repr__(self):
        return '%s @ %s km with %s points' % (self.name, self.position, self.score)


def read_input(filename):
    reindeer = []
    with open(filename, 'r') as f:
        for line in f:
            words = line.strip('.\n').split(' ')
            reindeer.append(Reindeer(
                name=words[0],
                speed=int(words[3]),
                flight_time=int(words[6]),
                rest_time=int(words[13])
            ))
    return reindeer
