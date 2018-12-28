package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/campoy/advent-of-code-2018/aoc"
)

const (
	steps       = 50000000000
	patternSize = 5
)

func main() {
	state, rules, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// speed at which the repeated state moves to the right.
	var speed int

	step := 0
	for ; step < steps; step++ {
		fmt.Printf("%5d %s\n", step, state)
		n := state.apply(rules)
		diff, ok := state.equals(n)
		state = n
		if ok {
			fmt.Println("found equilibrium with diff", diff)
			speed = diff
			break
		}
	}

	if remainingSteps := steps - step - 1; remainingSteps > 0 {
		fmt.Printf("... simulating %d remaining steps\n", remainingSteps)
		state = state.move(remainingSteps * speed)
		step += remainingSteps + 1
	}

	fmt.Printf("%5d %s\n", step, state)
	fmt.Println(state.sum())
}

func (s state) apply(rs rules) state {
	n := make(state)
	min, max := s.bounds()
	for i := min - patternSize; i <= max+patternSize; i++ {
		p := s.patternAt(i)
		if rs[p] {
			n[i] = true
		}
	}
	return n
}

func (s state) equals(t state) (int, bool) {
	min, max := s.bounds()
	tMin, tMax := t.bounds()
	if max-min != tMax-tMin {
		return 0, false
	}

	diff := tMin - min
	for i := min; i <= max; i++ {
		if s[i] != t[i+diff] {
			return 0, false
		}
	}

	return diff, true
}

func (s state) sum() int {
	sum := 0
	for i, p := range s {
		if p {
			sum += i
		}
	}
	return sum
}

func (s state) move(p int) state {
	n := make(state)
	for k, v := range s {
		n[k+p] = v
	}
	return n
}

type pot bool

func readPot(r rune) pot { return pot(r == '#') }

func (p pot) String() string {
	if p {
		return "#"
	}
	return "."
}

type pots []pot

func (ps pots) String() string {
	w := new(strings.Builder)
	for _, p := range ps {
		fmt.Fprint(w, p)
	}
	return w.String()
}

type state map[int]pot

func readState(text string) state {
	s := make(state, len(text))
	for i, c := range text {
		s[i] = readPot(c)
	}
	return s
}

func (s state) bounds() (int, int) {
	min, max := aoc.MaxInt, aoc.MinInt
	for p := range s {
		if s[p] {
			min = aoc.Min(min, p)
			max = aoc.Max(max, p)
		}
	}
	return min, max
}

func (s state) patternAt(i int) pattern {
	return pattern{s[i-2], s[i-1], s[i], s[i+1], s[i+2]}
}

func (s state) String() string {
	min, max := s.bounds()
	ps := make(pots, max-min+1)
	for i := min; i <= max; i++ {
		ps[i-min] = s[i]
	}
	return ps.String()
}

type pattern [patternSize]pot

func (p pattern) String() string { return pots(p[:]).String() }

type rules map[pattern]pot

func (r rules) readRule(text string) error {
	var left, right string
	_, err := fmt.Sscanf(text, "%s => %s", &left, &right)
	if err != nil {
		return fmt.Errorf("could not parse rule %q", text)
	}

	var p pattern
	for i, c := range left {
		p[i] = readPot(c)
	}
	r[p] = readPot(rune(right[0]))
	return nil
}

func (r rules) String() string {
	w := new(strings.Builder)
	for k, v := range r {
		fmt.Printf("%v => %v\n", pots(k[:]), v)
	}
	return w.String()
}

func readInput(path string) (state, rules, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	if !s.Scan() {
		return nil, nil, fmt.Errorf("input is too short")
	}
	state := readState(strings.TrimPrefix(s.Text(), "initial state: "))

	if !s.Scan() {
		return nil, nil, fmt.Errorf("input is too short")
	}
	rules := make(rules)
	for s.Scan() {
		if err := rules.readRule(s.Text()); err != nil {
			return nil, nil, err
		}
	}
	return state, rules, s.Err()
}
