package main

import (
	"os"
	"math/rand"
	"fmt"
	"github.com/deckarep/golang-set"
	"io/ioutil"
	"strings"
	"runtime/pprof"
)

type Replacement struct {
	long  string
	cost  int
	short string
}

func allIndices(s string, substring string) []int {
	offset := 0
	indices := []int{}
	for {
		index := strings.Index(s[offset:], substring)
		if index == -1 {
			return indices
		}
		indices = append(indices, offset+index)
		offset += index + 1
	}
	// shuffle the indices hoping for betterer results
	for i := range indices {
		j := rand.Intn(i + 1)
		indices[i], indices[j] = indices[j], indices[i]
	}
	return indices
}

var seen = mapset.NewSet()

func contract(molecule string, replacements []Replacement) Replacement {
	//fmt.Printf("> %s\n", molecule)
	if molecule == "e" {
		return Replacement{molecule, 0, molecule}
	}
	retval := Replacement{molecule, -1, molecule}
	if strings.Contains(molecule, "e") || seen.Contains(molecule) {
		return retval
	}
	seen.Add(molecule)

	// shuffle the replacements hoping for betterer results
	for i := range replacements {
		j := rand.Intn(i + 1)
		replacements[i], replacements[j] = replacements[j], replacements[i]
	}
	for _, replacement := range replacements {
		for _, start := range allIndices(molecule, replacement.long) {
			pre := molecule[:start]
			post := molecule[start+len(replacement.long):]
			new := pre + replacement.short + post

			rest := contract(new, replacements)
			if rest.cost < len(molecule) {
				return Replacement{molecule, rest.cost + replacement.cost, rest.short}
			}
		}
	}

	fmt.Printf("%d\t%d\t%s\n", seen.Cardinality(), len(molecule), molecule)
	return retval
}

func readInput(filename string) (string, []Replacement) {
	txt, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(txt), "\n")
	inputMolecule := lines[len(lines)-1]
	replacements := []Replacement{}
	for _, line := range lines[:len(lines)-2] {
		item := strings.Split(line, " => ")
		replacements = append(replacements, Replacement{item[1], 1, item[0]})
	}
	return inputMolecule, replacements
}

func main() {
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	molecule, replacements := readInput("input")
	//fmt.Println(replacements)
	fmt.Print(contract(molecule, replacements))
}
