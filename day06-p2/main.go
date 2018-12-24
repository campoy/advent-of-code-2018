package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	b, err := readBoard("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(explore(b, []pos{b.center()}, 10000))
}

type pos struct{ x, y int }
type board map[pos]int

func readBoard(path string) (board, error) {
	board := make(board)

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		var a, b int
		fmt.Sscanf(s.Text(), "%d, %d", &a, &b)
		board[pos{a, b}] = len(board)
	}
	return board, s.Err()
}

func (b board) center() pos {
	x, y := 0, 0
	for p := range b {
		x += p.x
		y += p.y
	}
	return pos{x / len(b), y / len(b)}
}

func (b board) sumOfDist(c pos) int {
	d := 0
	for p := range b {
		d += dist(c, p)
	}
	return d
}

func dist(p, q pos) int { return abs(p.x-q.x) + abs(p.y-q.y) }
func abs(v int) int     { return int(math.Abs(float64(v))) }

func explore(b board, ps []pos, d int) int {
	seen := make(map[pos]bool)
	area := 0

	for len(ps) > 0 {
		p := ps[0]
		ps = ps[1:]

		if seen[p] {
			continue
		}
		seen[p] = true

		if b.sumOfDist(p) >= d {
			continue
		}
		area++

		for _, n := range neighbors(p) {
			if seen[n] {
				continue
			}
			ps = append(ps, n)
		}
	}
	return area
}

func neighbors(p pos) []pos {
	return []pos{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y - 1}, {p.x, p.y + 1}}
}
