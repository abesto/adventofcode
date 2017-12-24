package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Maths taken straight from https://www.redblobgames.com/grids/hexagons/
// Using cube coordinates
type position struct {
	x int
	y int
	z int
}

type vector position

func mkPosition() *position {
	return &position{}
}

var directions = map[string]vector{
	"n":  vector{0, 1, -1},
	"ne": vector{1, 0, -1},
	"se": vector{1, -1, 0},
	"s":  vector{0, -1, 1},
	"sw": vector{-1, 0, 1},
	"nw": vector{-1, 1, 0},
}

func (p *position) moveVector(v vector) {
	p.x, p.y, p.z = p.x+v.x, p.y+v.y, p.z+v.z
}

func (p *position) move(direction string) {
	p.moveVector(directions[direction])
}

func absInt(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func distance(a *position, b *position) int {
	return (absInt(a.x-b.x) + absInt(a.y-b.y) + absInt(a.z-b.z)) / 2
}

func solve(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file)
	maxDist := 0
	origin := mkPosition()
	for reader.Scan() {
		p := mkPosition()
		for _, direction := range strings.Split(reader.Text(), ",") {
			p.move(direction)
			maxDist = maxInt(maxDist, distance(p, origin))
		}
		fmt.Printf("%d %d %d\n", p, distance(p, origin), maxDist)
	}
}

func main() {
	solve("input.sample1")
	solve("input.sample2")
	solve("input.sample3")
	solve("input.sample4")
	solve("input")
}
