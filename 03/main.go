package main

import (
	"fmt"
	"math"
)

/*
Math for part 1:
ringLength(n) = 8n
ringEnd(n) = sum(i:1->n)ringLength(n) = 1*8 + 2*8 + ... * n*8 = 8*sum(i:1->n)i = 8*n*(n-1)/2
x=8*n*(n-1)/2 -> 0 = n^2-n-x/4 -> ring(x) = floor((1+sqrt(1+x))/2)
posOnRing(x) = x - ringEnd(ring(x)-1)
*/

func ring(x int) int {
	return int((1 + math.Sqrt(1+float64(x))) / 2)
}

func ringEnd(n int) int {
	return (8 * n * (n - 1)) / 2
}

func ringLength(n int) int {
	return 8 * n
}

func posOnRing(x int) int {
	return x - ringEnd(ring(x))
}

func posOnRingWithKnownRing(x int, xRing int) int {
	return x - ringEnd(xRing)
}

func intDist(a int, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func steps(x int) int {
	xRing := ring(x)
	xPos := posOnRingWithKnownRing(x, xRing)
	xRingLength := ringLength(xRing)
	ringEigth := xRingLength / 8
	eigthPos := xPos / ringEigth
	if eigthPos%2 == 1 {
		return xRing - 1 + intDist(xPos, eigthPos*ringEigth)
	} else {
		return xRing - 1 + intDist(xPos, (eigthPos-1)*ringEigth)
	}
}

// Part two: not sure if there's a closed formula here. Just gonna brute-force it.
// Not even gonna bother to map the coordinates into a single array.

type vector struct {
	x int
	y int
}

func sumNeighbors(mem [100][100]int, x int, y int) int {
	var sum = 0
	for _, x := range []int{x - 1, x, x + 1} {
		for _, y := range []int{y - 1, y, y + 1} {
			sum += mem[x][y]
		}
	}
	return sum
}

var directions = [4]vector{vector{-1, 0}, vector{0, -1}, vector{1, 0}, vector{0, 1}}

func turnLeft(v vector) vector {
	for i := 0; i < 4; i++ {
		if directions[i] == v {
			return directions[(i+1)%4]
		}
	}
	panic("math has failed")
}

func solveTwo(limit int) int {
	var mem [100][100]int // That should be enough
	var x, y = 50, 50
	mem[x][y] = 1
	x += 1
	var direction = vector{0, 1}
	for {
		mem[x][y] = sumNeighbors(mem, x, y)
		//fmt.Printf("%d\t%d\t%d\n", x, y, mem[x][y])
		if mem[x][y] > limit {
			return mem[x][y]
		}
		leftVector := turnLeft(direction)
		leftX, leftY := x+leftVector.x, y+leftVector.y
		if mem[leftX][leftY] == 0 {
			direction = leftVector
			x, y = leftX, leftY
		} else {
			x, y = x+direction.x, y+direction.y
		}
	}
}

func main() {
	/*
		for i := 1; i < 49; i++ {
			fmt.Printf("i=%d ring(i)=%d posOnRing(i)=%d dist(i)=%d\n", i, ring(i), posOnRing(i), steps(i))
		}
	*/
	fmt.Printf("Sample 1\n1 -> %d\n12 -> %d\n23 -> %d\n1024 -> %d\n", steps(1), steps(12), steps(23), steps(1024))
	fmt.Printf("Solution 1: %d\n", steps(277678))
	fmt.Printf("Sample for 2\n132 -> %d (should be 133)\n", solveTwo(132))
	fmt.Printf("Solution 2: %d\n", solveTwo(277678))
}
