package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type lineSolver interface {
	Reset()
	Number(int)
	Solution() int
}

type lineSolverOne struct {
	min int
	max int
}

func (s *lineSolverOne) Reset() {
	s.min = int(^uint(0) >> 1)
	s.max = 0
}

func (s *lineSolverOne) Number(num int) {
	if num < s.min {
		s.min = num
	}
	if num > s.max {
		s.max = num
	}
}

func (s *lineSolverOne) Solution() int {
	diff := s.max - s.min
	fmt.Printf("min=%d max=%d diff=%d ", s.min, s.max, diff)
	return diff
}

func (s *lineSolverOne) String() string {
	return "lineSolverOne"
}

type lineSolverTwo struct {
	nums     []int
	done     bool
	solution int
}

func (s *lineSolverTwo) Reset() {
	s.done = false
	if s.nums == nil {
		s.nums = []int{}
	} else {
		s.nums = make([]int, 0, len(s.nums))
	}
}

func sortedPair(a int, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a
}

func (s *lineSolverTwo) Number(num int) {
	if s.done {
		return
	}
	for i := 0; i < len(s.nums); i++ {
		other := s.nums[i]
		a, b := sortedPair(num, other)
		if a%b == 0 {
			s.done = true
			s.solution = a / b
			fmt.Printf("%d %d sum=%d ", other, num, s.solution)
			return
		}
	}
	s.nums = append(s.nums, num)
}

func (s *lineSolverTwo) Solution() int {
	if !s.done {
		panic("OMGWTF")
	}
	return s.solution
}

func (s *lineSolverTwo) String() string {
	return "lineSolverTwo"
}

func solve(filename string, solver lineSolver) {
	fmt.Printf("Solving %s with %s\n", filename, solver)
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	lineScanner := bufio.NewScanner(reader)
	var solution = 0
	for lineScanner.Scan() {
		solver.Reset()
		wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		for wordScanner.Scan() {
			num, err := strconv.Atoi(wordScanner.Text())
			if err != nil {
				panic(err)
			}
			solver.Number(num)
		}
		solution += solver.Solution()
		fmt.Printf("solution=%d\n", solution)
	}
	fmt.Printf("%d\n", solution)

}

func solveOne(filename string) {
}

func main() {
	solverOne := lineSolverOne{}
	solve("input.sample1", &solverOne)
	solve("input", &solverOne)
	solverTwo := lineSolverTwo{}
	solve("input.sample2", &solverTwo)
	solve("input", &solverTwo)
}
