package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	name     string
	weight   int
	parent   *node
	children []*node
}

func newNode(name string, weight int) *node {
	n := node{name, weight, nil, []*node{}}
	return &n
}

func (parent *node) addChild(child *node) {
	if child.parent != nil {
		log.Panicf("Parent of %s is already set to %s, refusing to set to %s", child.name, child.parent.name, parent.name)
	}
	child.parent = parent
	parent.children = append(parent.children, child)
}

func (n *node) subtreeWeight() int {
	s := n.weight
	for _, child := range n.children {
		s += child.subtreeWeight()
	}
	return s
}

func (n *node) isBalanced() bool {
	if len(n.children) == 0 {
		return true
	}
	weight := n.children[0].subtreeWeight()
	for i := 1; i < len(n.children); i++ {
		if n.children[i].subtreeWeight() != weight {
			return false
		}
	}
	return true
}

func (n *node) isRoot() bool {
	return n.parent == nil
}

func loadTree(filename string) *node {
	log.Printf("Loading tree from %s", filename)
	nodes := map[string]*node{}
	reader, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		wordScanner := bufio.NewScanner(strings.NewReader(scanner.Text()))
		wordScanner.Split(bufio.ScanWords)
		// Read the name
		wordScanner.Scan()
		name := wordScanner.Text()
		// Then the weight
		wordScanner.Scan()
		weightStr := wordScanner.Text()
		if weightStr[0] != '(' {
			panic("WTF (")
		}
		if weightStr[len(weightStr)-1] != ')' {
			panic("WTF )")
		}
		weight, err := strconv.Atoi(weightStr[1 : len(weightStr)-1])
		if err != nil {
			panic(err)
		}
		// Get the node object
		promise, exists := nodes[name]
		var n *node
		if exists {
			// If we've already read the parent of the node, then let's read the "promise" we created
			n = promise
		} else {
			// Otherwise we have a new node!
			n = newNode(name, weight)
			nodes[n.name] = n
		}
		n.weight = weight
		// Finally, handle the children, if there are any
		if wordScanner.Scan() {
			if wordScanner.Text() != "->" {
				log.Panicf("Expected to read '->', found '%s' instead", wordScanner.Text())
			}
			for wordScanner.Scan() {
				childName := wordScanner.Text()
				if childName[len(childName)-1] == ',' {
					childName = childName[0 : len(childName)-1]
				}
				child, childExists := nodes[childName]
				if !childExists {
					child = newNode(childName, -1)
					nodes[childName] = child
				}
				n.addChild(child)
			}
		}
	}
	// Find root, checking in the meanwhile that the tree is connected (has a single root)
	var root *node
	for _, n := range nodes {
		if n.isRoot() {
			if root != nil {
				log.Panicf("Have two roots: %s and %s", root.name, n.name)
			}
			root = n
		}
	}
	return root
}

func printTree(root *node, depth int) {
	for i := 0; i < depth*4; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%s\n", root.name)
	for _, child := range root.children {
		printTree(child, depth+1)
	}
}

func solveTwo(n *node) int {
	if n.isBalanced() {
		return 0
	}
	var unbalancedChild *node
	for _, child := range n.children {
		if !child.isBalanced() {
			if unbalancedChild != nil {
				log.Panicf("%s has at least two unbalanced children: %s and %s", n.name, unbalancedChild.name, child.name)
			}
			unbalancedChild = child
		}
	}
	if unbalancedChild == nil {
		// We've found the direct parent of the wrong node
		log.Printf("%s is the direct parent of the wrong node", n.name)
		weights := map[int][]*node{}
		for _, child := range n.children {
			weight := child.subtreeWeight()
			if _, exists := weights[weight]; !exists {
				weights[weight] = []*node{child}
			} else {
				weights[weight] = append(weights[weight], child)
			}
		}
		if len(weights) != 2 {
			panic("WTF >2 weights")
		}
		lonelyWeight := -1
		nonLonelyWeight := -1
		for weight, children := range weights {
			if len(children) == 1 {
				if lonelyWeight != -1 {
					panic("WTF >1 lonely weights")
				}
				lonelyWeight = weight
			} else if len(children) > 1 || nonLonelyWeight == -1 {
				nonLonelyWeight = weight
			}
		}
		return weights[lonelyWeight][0].weight + nonLonelyWeight - lonelyWeight
	} else {
		// Recurse!
		return solveTwo(unbalancedChild)
	}
}

func main() {
	sampleTree := loadTree("input.sample")
	tree := loadTree("input")
	log.Printf("Root of full input: %s\n", tree.name)
	log.Printf("Weight correction for sample: %d\n", solveTwo(sampleTree))
	log.Printf("Weight correction for actual input: %d\n", solveTwo(tree))
}
