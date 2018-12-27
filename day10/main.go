package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/campoy/advent-of-code-2018/aoc"
)

func main() {
	ps, err := readParticles("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	dx, dy := ps.size()
	for i := 0; ; i++ {
		ps.step()
		ndx, ndy := ps.size()
		if ndx > dx || ndy > dy {
			ps.reverse()
			fmt.Printf("After waiting for %d seconds you read:\n\n", i)
			break
		}
		dx, dy = ndx, ndy
	}
	fmt.Println(ps)
}

type vector struct{ x, y int }
type particle struct{ p, v vector }
type particles struct {
	min, max vector
	ps       []particle
}

func readParticles(path string) (*particles, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var ps particles
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y, vx, vy int
		_, err := fmt.Sscanf(s.Text(), "position=<%d,%d> velocity=<%d,%d>", &x, &y, &vx, &vy)
		if err != nil {
			return nil, fmt.Errorf("couldn't parse %q: %v", s.Text(), err)
		}
		if len(ps.ps) == 0 {
			ps.min.x, ps.min.y = x, y
			ps.max.x, ps.max.y = x, y
		} else {
			ps.min.x, ps.min.y = aoc.Min(ps.min.x, x), aoc.Min(ps.min.y, y)
			ps.max.x, ps.max.y = aoc.Max(ps.max.x, x), aoc.Max(ps.max.y, y)
		}
		ps.ps = append(ps.ps, particle{vector{x, y}, vector{vx, vy}})
	}
	return &ps, s.Err()
}

func (ps *particles) size() (dx, dy int) {
	dx = ps.max.x - ps.min.x + 1
	dy = ps.max.y - ps.min.y + 1
	return dx, dy
}

func (ps *particles) String() string {
	m := make(map[vector]bool)
	for _, p := range ps.ps {
		m[p.p] = true
	}
	w := new(strings.Builder)
	for y := ps.min.y; y <= ps.max.y; y++ {
		for x := ps.min.x; x <= ps.max.x; x++ {
			if m[vector{x, y}] {
				fmt.Fprint(w, "#")
			} else {
				fmt.Fprint(w, " ")
			}
		}
		fmt.Fprintln(w)
	}
	return w.String()
}

func (ps *particles) step() {
	ps.min.x, ps.min.y = ps.ps[0].p.x, ps.ps[0].p.y
	ps.max.x, ps.max.y = ps.ps[0].p.x, ps.ps[0].p.y
	for i := range ps.ps {
		ps.ps[i].p.x += ps.ps[i].v.x
		ps.ps[i].p.y += ps.ps[i].v.y
		ps.min.x, ps.min.y = aoc.Min(ps.min.x, ps.ps[i].p.x), aoc.Min(ps.min.y, ps.ps[i].p.y)
		ps.max.x, ps.max.y = aoc.Max(ps.max.x, ps.ps[i].p.x), aoc.Max(ps.max.y, ps.ps[i].p.y)
	}
}

func (ps *particles) reverse() {
	p0 := ps.ps[0].p
	ps.min.x, ps.min.y = p0.x, p0.y
	ps.max.x, ps.max.y = p0.x, p0.y
	for i := range ps.ps {
		ps.ps[i].p.x -= ps.ps[i].v.x
		ps.ps[i].p.y -= ps.ps[i].v.y
		ps.min.x, ps.min.y = aoc.Min(ps.min.x, ps.ps[i].p.x), aoc.Min(ps.min.y, ps.ps[i].p.y)
		ps.max.x, ps.max.y = aoc.Max(ps.max.x, ps.ps[i].p.x), aoc.Max(ps.max.y, ps.ps[i].p.y)
	}
}
