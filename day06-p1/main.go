package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	b, err := readBoard("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)

	maxArea := 0
	var maxPos pos
	for _, f := range b.finitePoints() {
		a := b.area(f)
		if a > maxArea {
			maxArea = a
			maxPos = f
		}
	}

	fmt.Println(maxPos, maxArea)
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

func (b board) maxPos() pos {
	maxX, maxY := 0, 0

	for p := range b {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return pos{maxX, maxY}
}

func (b board) String() string {
	max := b.maxPos()

	var out strings.Builder
	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			id, ok := b[pos{x, y}]
			if !ok {
				if c := b.closest(pos{x, y}); c >= 0 {
					fmt.Fprintf(&out, string('A'+byte(c)))
				} else {
					fmt.Fprintf(&out, ".")
				}
				continue
			}
			fmt.Fprint(&out, string('A'+byte(id)))
		}
		fmt.Fprint(&out, "\n")
	}
	return out.String()
}

func (b board) closest(p pos) int {
	var ids []int
	min := -1
	for q := range b {
		d := dist(p, q)
		if min < 0 || min > d {
			min = d
			ids = []int{b[q]}
			continue
		}

		if min == d {
			ids = append(ids, b[q])
		}
	}

	if len(ids) == 1 {
		return ids[0]
	}
	return -1
}

func dist(p, q pos) int {
	return int(math.Abs(float64(p.x-q.x)) + math.Abs(float64(p.y-q.y)))
}

func (b board) finitePoints() []pos {
	var fs []pos
	for p := range b {
		var xy, XY, xY, Xy bool
		for q := range b {
			switch {
			case p.x > q.x && p.y > q.y:
				xy = true
			case p.x > q.x && p.y < q.y:
				xY = true
			case p.x < q.x && p.y > q.y:
				Xy = true
			case p.x < q.x && p.y < q.y:
				XY = true
			}
		}
		if xy && xY && Xy && XY {
			fs = append(fs, p)
		}
	}
	return fs
}

func (b board) area(p pos) int {
	max := b.maxPos()
	a := 0
	for x := 0; x <= max.x; x++ {
		for y := 0; y <= max.y; y++ {
			c := b.closest(pos{x, y})
			if c == b[p] {
				a++
			}
		}
	}
	return a
}
