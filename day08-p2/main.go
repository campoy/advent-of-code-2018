package main

import (
	"fmt"
	"log"

	"github.com/campoy/advent-of-code-2018/aoc"
)

func main() {
	nums, err := aoc.ReadIntSeqFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	n := readNode(nums)
	fmt.Println(n.value())
}

type node struct {
	children []node
	metadata []int
}

func readNode(nums *aoc.IntSeq) node {
	n := node{
		children: make([]node, nums.Next()),
		metadata: make([]int, nums.Next()),
	}
	for i := range n.children {
		n.children[i] = readNode(nums)
	}
	for i := range n.metadata {
		n.metadata[i] = nums.Next()
	}
	return n
}

func (n node) value() int {
	if len(n.children) == 0 {
		return aoc.Sum(n.metadata)
	}
	sum := 0
	for _, m := range n.metadata {
		if m == 0 || m > len(n.children) {
			continue
		}
		sum += n.children[m-1].value()
	}
	return sum
}
