package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type mutate func(int) int

func plusOne(x int) int {
	return x + 1
}

func decreaseThreeOrAboveOrIncrease(x int) int {
	if x >= 3 {
		return x - 1
	}
	return x + 1
}

func solve(filename string, mutate mutate) {
	fmt.Printf("Solving part one for %s: ", filename)
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	instructions := []int{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		offset, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, offset)
	}
	currentOffset := 0
	steps := 0
	for currentOffset >= 0 && currentOffset < len(instructions) {
		newOffset := currentOffset + instructions[currentOffset]
		instructions[currentOffset] = mutate(instructions[currentOffset])
		currentOffset = newOffset
		steps += 1
	}
	fmt.Printf("%d\n", steps)
}

func main() {
	solve("input.sample", plusOne)
	solve("input", plusOne)
	solve("input.sample", decreaseThreeOrAboveOrIncrease)
	solve("input", decreaseThreeOrAboveOrIncrease)
}
