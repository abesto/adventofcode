package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type securityCheck func(word string, seen map[string]bool) bool

func same(word string, seen map[string]bool) bool {
	_, seenThis := seen[word]
	return !seenThis
}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	l := len(a)
	matched := make([]bool, l)
	for ai := 0; ai < l; ai++ {
		var bi int
		for bi = 0; bi < l; bi++ {
			if a[ai] == b[bi] && !matched[bi] {
				matched[bi] = true
				break
			}
		}
		if bi == l {
			return false
		}
	}
	for i := 0; i < l; i++ {
		if !matched[i] {
			return false
		}
	}
	return true
}

func anagram(word string, seen map[string]bool) bool {
	for seenWord, _ := range seen {
		if isAnagram(word, seenWord) {
			return false
		}
	}
	return true
}

func isValid(passphrase string, check securityCheck) bool {
	seen := make(map[string]bool)
	scanner := bufio.NewScanner(strings.NewReader(passphrase))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if !check(word, seen) {
			return false
		} else {
			seen[word] = true
		}
	}
	return true
}

func solve(check securityCheck) int {
	reader, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	lineScanner := bufio.NewScanner(reader)
	var solution = 0
	for lineScanner.Scan() {
		if isValid(lineScanner.Text(), check) {
			solution += 1
		}
	}
	return solution
}

func main() {
	fmt.Printf("Samples for part one:\naa bb cc dd ee: %t\naa bb cc dd aa: %t\naa bb cc dd aaa: %t\n",
		isValid("aa bb cc dd ee", same), isValid("aa bb cc dd aa", same), isValid("aa bb cc dd aaa", same))
	fmt.Printf("Solution for part one: %d\n", solve(same))
	fmt.Printf("Samples for part two:\n")
	for _, phrase := range []string{"abcde fghij", "abcde xyz ecdab", "a ab abc abd abf abj", "iiii oiii ooii oooi oooo", "oiii ioii iioi iiio"} {
		fmt.Printf("%s: %t\n", phrase, isValid(phrase, anagram))
	}
	fmt.Printf("Solution for part one: %d\n", solve(anagram))
}
