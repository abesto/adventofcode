package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	id        int
	graph     *graph
	neighbors []*vertex
}

func (root *vertex) computeGroup() map[*vertex]bool {
	visited := map[*vertex]bool{}
	queue := []*vertex{root}
	var v *vertex
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		if visited[v] {
			continue
		}
		visited[v] = true
		queue = append(queue, v.neighbors...)
	}
	return visited

}

func connect(a *vertex, b *vertex) {
	a.neighbors = append(a.neighbors, b)
	b.neighbors = append(b.neighbors, a)
}

type graph struct {
	vertices map[int]*vertex
}

func mkGraph() *graph {
	return &graph{map[int]*vertex{}}
}

func (g *graph) countGroups() int {
	count := 0
	visited := map[*vertex]bool{}
	for _, root := range g.vertices {
		if visited[root] {
			continue
		}
		count += 1
		for v, _ := range root.computeGroup() {
			visited[v] = true
		}
	}
	return count
}

func (g *graph) vertex(id int) *vertex {
	v, exists := g.vertices[id]
	if exists {
		return v
	}
	v = &vertex{id, g, []*vertex{}}
	g.vertices[id] = v
	return v
}

func readGraph(filename string) *graph {
	g := mkGraph()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		// Read the node ID
		wordScanner.Scan()
		id, err := strconv.Atoi(wordScanner.Text())
		lhs := g.vertex(id)
		if err != nil {
			panic(err)
		}
		// Next, read <->
		wordScanner.Scan()
		if wordScanner.Text() != "<->" {
			panic("WTF, expected '<->', got: " + wordScanner.Text())
		}
		// Next, read the connected vertices
		for wordScanner.Scan() {
			idStr := wordScanner.Text()
			if idStr[len(idStr)-1] == ',' {
				idStr = idStr[0 : len(idStr)-1]
			}
			id, err := strconv.Atoi(idStr)
			if err != nil {
				panic(err)
			}
			rhs := g.vertex(id)
			connect(lhs, rhs)
		}
	}
	return g
}

func solve(filename string) {
	g := readGraph(filename)
	fmt.Printf("%d %d\n",
		len(g.vertex(0).computeGroup()),
		g.countGroups(),
	)
}

func solveTwo(filename string) {
}

func main() {
	solve("input.sample")
	solve("input")
}
