package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/campoy/advent-of-code-2018/aoc"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var players, last int
	_, err = fmt.Sscanf(string(b), "%d players; last marble is worth %d points", &players, &last)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(score(players, last, nil))
}

type ring struct {
	value       int
	right, left *ring
}

func (r *ring) add(v int) (*ring, int) {
	if r == nil {
		n := &ring{value: 0}
		n.right = n
		n.left = n
		return n, 0
	}

	if v%23 != 0 {
		a := r.right
		b := a.right
		n := &ring{value: v, left: a, right: b}
		a.right = n
		b.left = n
		return n, 0
	}

	drop := r
	for i := 0; i < 7; i++ {
		drop = drop.left
	}
	drop.left.right = drop.right
	drop.right.left = drop.left
	return drop.right, v + drop.value
}

func score(players, last int, ring *ring) int {
	ps := make([]int, players)
	for i := 0; i <= last; i++ {
		next, value := ring.add(i)
		ps[i%players] += value
		ring = next
	}
	return aoc.Max(ps...)
}
