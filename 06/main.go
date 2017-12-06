package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func iterate(banks []int) {
	biggestBankIndex := 0
	for i := 1; i < len(banks); i++ {
		if banks[i] > banks[biggestBankIndex] {
			biggestBankIndex = i
		}
	}
	// Inputs are small, not gonna optimize out the iterations here
	i := biggestBankIndex
	blocks := banks[biggestBankIndex]
	banks[biggestBankIndex] = 0
	for blocks > 0 {
		i = (i + 1) % len(banks)
		banks[i] += 1
		blocks -= 1
	}
}

func banksEqual(a []int, b []int) bool {
	for i := 0; i < len(a); i += 1 {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func seenBankConfiguration(banks []int, seen [][]int) bool {
	for i := 0; i < len(seen); i++ {
		var j int
		for j = 0; j < len(banks); j++ {
			if banks[j] != seen[i][j] {
				break
			}
		}
		if j == len(banks) {
			return true
		}
	}
	return false
}

func printBanks(banks []int) {
	for i := 0; i < len(banks); i++ {
		fmt.Printf("%d ", banks[i])
	}
	fmt.Println()
}

func solve(filename string) {
	banks := []int{}
	fmt.Printf("Solving for %s: ", filename)
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		blocks, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		banks = append(banks, blocks)
	}
	seen := [][]int{}
	for !seenBankConfiguration(banks, seen) {
		snapshot := make([]int, len(banks))
		copy(snapshot, banks)
		seen = append(seen, snapshot)
		iterate(banks)
	}
	fmt.Printf("%d ", len(seen))
	pattern := make([]int, len(banks))
	copy(pattern, banks)
	cycles := 1
	for iterate(banks); !banksEqual(banks, pattern); cycles += 1 {
		iterate(banks)
	}
	fmt.Printf("%d\n", cycles)
}

func main() {
	solve("input.sample")
	solve("input")
}
