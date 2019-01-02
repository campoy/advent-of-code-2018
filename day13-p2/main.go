package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/campoy/advent-of-code-2018/aoc"
)

func main() {
	debug := flag.Bool("v", false, "verbose")
	flag.Parse()

	plan, err := readPlan("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for ; len(plan.carts) > 1; plan.tick() {
		if *debug {
			// fmt.Printf("\033[H\033[2J%v", plan)
			fmt.Println(plan)
			time.Sleep(time.Second)
		}
	}
	fmt.Println(plan.carts[0].pos)
}

type cell byte

const (
	empty      cell = ' '
	horizontal cell = '-'
	vertical   cell = '|'
	turnL      cell = '\\'
	turnJ      cell = '/'
	crossing   cell = '+'
)

func isCell(b rune) bool {
	switch cell(b) {
	case empty, horizontal, vertical, turnL, turnJ, crossing:
		return true
	default:
		return false
	}
}

func (c cell) String() string { return string(c) }

func (c cell) links(d direction) bool {
	if d == up || d == down {
		return c != empty && c != horizontal
	}
	return c != empty && c != vertical
}

type direction byte

const (
	up direction = iota
	right
	down
	left
)

var (
	nameToDir = map[rune]direction{'^': up, '>': right, 'v': down, '<': left}
	dirToName = map[direction]string{up: "^", right: ">", down: "v", left: "<"}
	dirToVec  = map[direction]pos{up: {0, -1}, right: {1, 0}, down: {0, 1}, left: {-1, 0}}
)

func readDirection(r rune) direction     { return nameToDir[r] }
func (d direction) String() string       { return dirToName[d] }
func (d direction) turnLeft() direction  { return direction((d - 1) % 4) }
func (d direction) turnRight() direction { return direction((d + 1) % 4) }

type pos struct{ x, y int }

func (p pos) move(d direction) pos {
	v := dirToVec[d]
	return pos{p.x + v.x, p.y + v.y}
}

type cart struct {
	pos
	dir       direction
	crossings int
}

type plan struct {
	width, height int
	carts         []*cart
	cells         map[pos]cell
}

func readPlan(path string) (*plan, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("could not open input.txt: %v", err)
	}
	defer f.Close()

	plan := &plan{cells: make(map[pos]cell)}

	s := bufio.NewScanner(f)
	for y := 0; s.Scan(); y++ {
		text := s.Text()
		for x, c := range text {
			if isCell(c) {
				plan.cells[pos{x, y}] = cell(c)
				continue
			}
			cart := cart{
				pos: pos{x, y},
				dir: readDirection(c),
			}
			switch cart.dir {
			case up, down:
				plan.cells[pos{x, y}] = vertical
			case left, right:
				plan.cells[pos{x, y}] = horizontal
			}
			plan.carts = append(plan.carts, &cart)
		}
		plan.height++
		plan.width = aoc.Max(plan.width, len(text))
	}
	return plan, s.Err()
}

func (p *plan) cell(i, j int) cell {
	if c, ok := p.cells[pos{i, j}]; ok {
		return c
	}
	return empty
}

func (p *plan) String() string {
	w := new(strings.Builder)
	cartPos := make(map[pos]*cart)
	for _, c := range p.carts {
		cartPos[c.pos] = c
	}
	for y := 0; y < p.height; y++ {
		for x := 0; x < p.width; x++ {
			if c, ok := cartPos[pos{x, y}]; ok {
				fmt.Fprint(w, c.dir)
			} else {
				fmt.Fprintf(w, "%c", p.cell(x, y))
			}
		}
		fmt.Fprintln(w)
	}
	return w.String()
}

func (p *plan) tick() {
	sort.Slice(p.carts, func(i, j int) bool {
		if p.carts[i].y == p.carts[j].y {
			return p.carts[i].x < p.carts[j].x
		}
		return p.carts[i].y < p.carts[j].y
	})

	busy := make(map[pos]*cart)
	for _, c := range p.carts {
		busy[c.pos] = c
	}

	for cartIndex := 0; cartIndex < len(p.carts); cartIndex++ {
		c := p.carts[cartIndex]
		delete(busy, c.pos)
		c.pos = c.pos.move(c.dir)
		switch p.cells[c.pos] {
		case turnL:
			switch c.dir {
			case up:
				c.dir = left
			case right:
				c.dir = down
			case down:
				c.dir = right
			case left:
				c.dir = up
			}
		case turnJ:
			switch c.dir {
			case up:
				c.dir = right
			case right:
				c.dir = up
			case down:
				c.dir = left
			case left:
				c.dir = down
			}
		case crossing:
			switch c.crossings % 3 {
			case 0:
				c.dir = c.dir.turnLeft()
			case 1:
				// go straight
			case 2:
				c.dir = c.dir.turnRight()
			}
			c.crossings++
		}
		prev, ok := busy[c.pos]
		if !ok {
			busy[c.pos] = c
			continue
		}

		// remove the two carts
		fmt.Printf("crash detected at (%d,%d) removing both carts\n", c.x, c.y)
		delete(busy, c.pos)
		for i := 0; i < len(p.carts); i++ {
			if p.carts[i] == prev || p.carts[i] == c {
				p.carts = append(p.carts[:i], p.carts[i+1:]...)
				if i <= cartIndex {
					cartIndex--
				}
				i--
			}
		}
	}
}
