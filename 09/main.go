package main

import (
	"bufio"
	"fmt"
	"os"
)

type transition struct {
	deltaScore int
	nextState  *state
}

type state struct {
	name        string
	transitions map[byte]transition
}

func (s state) accept(c byte, deltaScore int, nextState *state) {
	s.transitions[c] = transition{deltaScore, nextState}
}

func (s state) transition(c byte) transition {
	transition, exists := s.transitions[c]
	if exists {
		return transition
	}
	defaultTransition, defaultExists := s.transitions[0]
	if defaultExists {
		return defaultTransition
	}
	panic("state " + s.name + " does not accept " + string(c))
}

type parser struct {
	states       map[string]*state
	currentState *state
	score        int
}

func mkParser() *parser {
	p := parser{map[string]*state{}, nil, 0}
	p.state("end")
	return &p
}

func (p *parser) step(c byte) {
	transition := p.currentState.transition(c)
	//fmt.Printf("(%s, %d) --%s, %d--> (%s, %d)\n", p.currentState.name, p.score, string(c), transition.deltaScore, transition.nextState.name, p.score+transition.deltaScore)
	p.currentState = transition.nextState
	p.score += transition.deltaScore
	if p.score == 0 {
		p.currentState = p.state("end")
	}
}

func (p *parser) state(name string) *state {
	s, exists := p.states[name]
	if !exists {
		s = &state{name, map[byte]transition{}}
		p.states[name] = s
	}
	return s
}

func buildParser() *parser {
	p := mkParser()
	start := p.state("start")
	groupStart := p.state("groupStart")
	groupEnd := p.state("groupEnd")
	comma := p.state("comma")
	garbage := p.state("garbage")
	garbageEnd := p.state("garbageEnd")
	bang := p.state("bang")
	skip := p.state("skip")

	start.accept('{', +1, groupStart)

	groupStart.accept('{', +1, groupStart)
	groupStart.accept('}', -1, groupEnd)
	groupStart.accept('<', 0, garbage)

	groupEnd.accept('}', -1, groupEnd)
	groupEnd.accept(',', 0, comma)

	comma.accept('{', +1, groupStart)
	comma.accept('<', 0, garbage)

	garbage.accept('!', 0, bang)
	garbage.accept('>', 0, garbageEnd)
	garbage.accept(0, 0, garbage)

	bang.accept(0, 0, skip)

	skip.accept('!', 0, bang)
	skip.accept('>', 0, garbageEnd)
	skip.accept(0, 0, garbage)

	garbageEnd.accept('}', -1, groupEnd)
	garbageEnd.accept(',', 0, comma)

	p.currentState = groupStart
	return p
}

func sumScore(reader *bufio.Reader) int {
	p := buildParser()
	sumScore := 0
	end := p.state("end")
	groupEnd := p.state("groupEnd")
	for p.currentState != end {
		c, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		p.step(c)
		if p.currentState == groupEnd {
			sumScore += p.score + 1
		}
	}
	sumScore += p.score + 1
	return sumScore
}

func countGarbage(reader *bufio.Reader) int {
	p := buildParser()
	count := 0
	end := p.state("end")
	garbage := p.state("garbage")
	garbageEnd := p.state("garbageEnd")
	for p.currentState != end {
		c, err := reader.ReadByte()
		if err != nil {
			panic(err)
		}
		p.step(c)
		if p.currentState == garbage {
			count += 1
		}
		if p.currentState == garbageEnd {
			count -= 1 // We count the opening '<', so reduce the result by one for each garbage string
		}
	}
	return count
}

func solve(filename string, solver func(reader *bufio.Reader) int) {
	fmt.Printf("Solving %s\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	for reader := bufio.NewReader(file); err == nil; _, err = reader.Discard(1) {
		fmt.Printf("%d\n", solver(reader))
	}
}

func main() {
	fmt.Printf("Part 1\n")
	solve("input.sample1", sumScore)
	solve("input", sumScore)
	fmt.Printf("Part 2\n")
	solve("input.sample2", countGarbage)
	solve("input", countGarbage)
}
