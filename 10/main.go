package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type hash struct {
	data     [256]byte
	position byte
	skip     int
}

func mkHash() *hash {
	hash := hash{}
	for i := 0; i < len(hash.data); i++ {
		hash.data[i] = byte(i)
	}
	return &hash
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func (hash *hash) twist(length byte) {
	items := make([]byte, length)
	for i := byte(0); i < length; i++ {
		items[length-1-i] = hash.data[hash.position+i]
	}
	for i := byte(0); i < length; i++ {
		hash.data[hash.position+i] = items[i]
	}
	hash.position += byte(int(length) + hash.skip)
	hash.skip += 1
}

// I love how Golang can figure out what part of speech each "hash" is,
// and doesn't complain about me overloading the hell out of the word.
func (hash *hash) hash() int {
	return int(hash.data[0]) * int(hash.data[1])
}

func solveOne(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		hash := mkHash()
		for _, lengthString := range strings.Split(reader.Text(), ",") {
			length, err := strconv.Atoi(lengthString)
			if err != nil {
				panic(err)
			}
			hash.twist(byte(length))
		}
		fmt.Println(hash.hash())
	}
}

var lengthSuffix = []byte{17, 31, 73, 47, 23}

func solveTwo(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		hash := mkHash()
		input := []byte(reader.Text())
		for i := 0; i < 64; i++ {
			for _, length := range input {
				hash.twist(length)
			}
			for _, length := range lengthSuffix {
				hash.twist(length)
			}
		}
		sparseHash := hash.data
		denseHash := [16]byte{}
		for i := 0; i < 16; i++ {
			denseHash[i] = sparseHash[16*i]
			for j := 1; j < 16; j++ {
				denseHash[i] ^= sparseHash[16*i+j]
			}
		}
		fmt.Printf("%x\n", denseHash)
	}
}

func main() {
	solveOne("input")
	solveTwo("input.sample2")
	solveTwo("input")
}
