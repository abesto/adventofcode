package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type packet struct {
	layer int
}

func (p *packet) advance() {
	p.layer += 1
}

func (p *packet) reset() {
	p.layer = -1
}

func mkPacket() *packet {
	return &packet{-1}
}

type layer struct {
	depth           int
	lrange          int
	scannerPosition int
	scannerSpeed    int
}

func (l *layer) advance() {
	if l.scannerPosition+l.scannerSpeed == l.lrange || l.scannerPosition+l.scannerSpeed == -1 {
		l.scannerSpeed *= -1
	}
	l.scannerPosition += l.scannerSpeed
}

func (l *layer) severity() int {
	return l.depth * l.lrange
}

func (l *layer) reset() {
	l.scannerPosition = 0
	l.scannerSpeed = 1
}

func mkLayer(depth int, lrange int) *layer {
	return &layer{depth, lrange, 0, 1}
}

type firewall struct {
	layers map[int]*layer
	depth  int
}

func (f *firewall) advance() {
	for _, layer := range f.layers {
		layer.advance()
	}
}

func (f *firewall) addLayer(depth int, lrange int) {
	f.layers[depth] = mkLayer(depth, lrange)
	if depth > f.depth {
		f.depth = depth
	}
}

func (f *firewall) reset() {
	for _, layer := range f.layers {
		layer.reset()
	}
}

func mkFirewall() *firewall {
	return &firewall{map[int]*layer{}, 0}
}

func (f1 *firewall) clone() *firewall {
	f2 := mkFirewall()
	f2.depth = f1.depth
	for _, l1 := range f1.layers {
		f2.layers[l1.depth] = &layer{l1.depth, l1.lrange, l1.scannerPosition, l1.scannerSpeed}
	}
	return f2
}

type simulation struct {
	firewall *firewall
	packet   *packet
	severity int
}

func (s *simulation) isDone() bool {
	return s.packet.layer == s.firewall.depth

}

func (s *simulation) isCaught() bool {
	currentLayer, exists := s.firewall.layers[s.packet.layer]
	if exists {
		//fmt.Printf("check: packet=%d depth=%d range=%d lsp=%d lseverity=%d severityBefore=%d\n", s.packet.layer, currentLayer.depth, currentLayer.lrange, currentLayer.scannerPosition, currentLayer.severity(), s.severity)

	}
	if exists && currentLayer.scannerPosition == 0 {
		//fmt.Printf("caught: depth=%d range=%d lseverity=%d severityBefore=%d\n", currentLayer.depth, currentLayer.lrange, currentLayer.severity(), s.severity)
		return true
	}
	return false
}

func (s *simulation) advance() {
	if s.isDone() {
		panic("done")
	}
	s.packet.advance()
	if s.isCaught() {
		s.severity += s.firewall.layers[s.packet.layer].severity()
	}
	s.firewall.advance()
}

func (s *simulation) run() *simulation {
	for !s.isDone() {
		s.advance()
	}
	return s
}

func (s *simulation) reset(f *firewall) {
	s.firewall = f
	s.packet.reset()
}

func (s *simulation) solveTwo() int {
	guess := 0
	for {
		initFirewall := s.firewall.clone()
		for !s.isDone() {
			s.packet.advance()
			if s.isCaught() {
				break
			}
			s.firewall.advance()
		}
		if !s.isCaught() {
			return guess
		}
		guess += 1
		initFirewall.advance()
		s.reset(initFirewall)
		//fmt.Println(guess)
	}
}

func mkSimulation() *simulation {
	s := simulation{
		mkFirewall(),
		mkPacket(),
		0,
	}
	return &s
}

func readFirewall(filename string) *simulation {
	s := mkSimulation()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		colonIndex := strings.Index(line, ":")
		depthStr, rangeStr := line[:colonIndex], line[colonIndex+2:]
		depth, err := strconv.Atoi(depthStr)
		if err != nil {
			panic(err)
		}
		lrange, err := strconv.Atoi(rangeStr)
		if err != nil {
			panic(err)
		}
		s.firewall.addLayer(depth, lrange)
	}
	return s
}

func main() {
	fmt.Println(readFirewall("input.sample").run().severity)
	fmt.Println(readFirewall("input").run().severity)
	fmt.Println(readFirewall("input.sample").solveTwo())
	fmt.Println(readFirewall("input").solveTwo())
}
