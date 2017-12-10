package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Modifiers: (int, int) -> int
type modifier func(lhs int, rhs int) int

func inc(lhs int, rhs int) int {
	return lhs + rhs
}

func dec(lhs int, rhs int) int {
	return lhs - rhs
}

// Operations: partially applied modifiers
type operation interface {
	evaluate(lhs int) int
}

type operationImpl struct {
	m   modifier
	rhs int
}

func (op *operationImpl) evaluate(lhs int) int {
	return op.m(lhs, op.rhs)
}

func mkOperation(wordScanner *bufio.Scanner) operation {
	wordScanner.Scan()
	opStr := wordScanner.Text()
	var m modifier
	if opStr == "inc" {
		m = inc
	} else if opStr == "dec" {
		m = dec
	} else {
		panic("WTF operation " + opStr)
	}
	wordScanner.Scan()
	rhs, err := strconv.Atoi(wordScanner.Text())
	if err != nil {
		panic(err)
	}
	return &operationImpl{m, rhs}
}

// Comparators: (int, int) -> bool
type comparator func(lhs int, rhs int) bool

func ge(lhs int, rhs int) bool {
	return lhs >= rhs
}

func gt(lhs int, rhs int) bool {
	return lhs > rhs
}

func eq(lhs int, rhs int) bool {
	return lhs == rhs
}

func lt(lhs int, rhs int) bool {
	return lhs < rhs
}

func le(lhs int, rhs int) bool {
	return lhs <= rhs
}

func ne(lhs int, rhs int) bool {
	return lhs != rhs
}

func mkComparator(input string) comparator {
	if input == ">=" {
		return ge
	}
	if input == ">" {
		return gt
	}
	if input == "==" {
		return eq
	}
	if input == "<" {
		return lt
	}
	if input == "<=" {
		return le
	}
	if input == "!=" {
		return ne
	}
	panic("WTF comparator " + input)
}

// Condition: partially applied operation
type condition interface {
	evaluate(cpu registers) bool
}

type conditionImpl struct {
	cmp comparator
	lhs string
	rhs int
}

func (c *conditionImpl) evaluate(cpu registers) bool {
	return c.cmp(cpu.read(c.lhs), c.rhs)
}

func mkCondition(wordScanner *bufio.Scanner) condition {
	wordScanner.Scan()
	if wordScanner.Text() != "if" {
		panic("Expected 'if', got " + wordScanner.Text())
	}
	wordScanner.Scan()
	lhs := wordScanner.Text()
	wordScanner.Scan()
	cmp := mkComparator(wordScanner.Text())
	wordScanner.Scan()
	rhs, err := strconv.Atoi(wordScanner.Text())
	if err != nil {
		panic(err)
	}
	return &conditionImpl{cmp, lhs, rhs}
}

// State (CPU registers)
type registers map[string]int

func (r registers) read(name string) int {
	val, exists := r[name]
	if exists {
		return val
	}
	return 0
}

// Instructions describe the input data fully
type instruction struct {
	reg  string
	op   operation
	cond condition
}

func (inst *instruction) apply(cpu registers) {
	old, exists := cpu[inst.reg]
	if !exists {
		old = 0
		cpu[inst.reg] = 0
	}
	if inst.cond.evaluate(cpu) {
		cpu[inst.reg] = inst.op.evaluate(old)
	}
}

func mkInstruction(line string) *instruction {
	wordScanner := bufio.NewScanner(strings.NewReader(line))
	wordScanner.Split(bufio.ScanWords)
	wordScanner.Scan()
	reg := wordScanner.Text()
	op := mkOperation(wordScanner)
	cond := mkCondition(wordScanner)
	return &instruction{reg, op, cond}
}

// Solution time!
func readInstructions(filename string) []*instruction {
	instructions := []*instruction{}
	return instructions
}

func findBiggestRegister(cpu registers) int {
	max := math.MinInt32
	for _, value := range cpu {
		if value > max {
			max = value
		}
	}
	return max
}

func solve(srcPath string) {
	fmt.Printf("Solving %s: ", srcPath)
	cpu := registers{}
	reader, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	maxAny := math.MinInt32
	for scanner.Scan() {
		mkInstruction(scanner.Text()).apply(cpu)
		maxCurrent := findBiggestRegister(cpu)
		if maxCurrent > maxAny {
			maxAny = maxCurrent
		}
	}
	fmt.Printf("maxEnd=%d maxAny=%d\n", findBiggestRegister(cpu), maxAny)
}

func main() {
	solve("input.sample")
	solve("input")
}
