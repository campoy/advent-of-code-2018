package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var fabric fabric
	s := bufio.NewScanner(f)
	for s.Scan() {
		var id, x, y, w, h int
		_, err := fmt.Sscanf(s.Text(), "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		if err != nil {
			log.Fatal(err)
		}
		fabric.addClaim(id, x, y, w, h)
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fabric.safeOption())
}

type xy struct{ x, y int }

type fabric struct {
	ids []int
	m   map[xy][]int
}

func (f *fabric) addClaim(id, x, y, w, h int) {
	if f.m == nil {
		f.m = make(map[xy][]int)
	}

	f.ids = append(f.ids, id)

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			p := xy{x + i, y + j}
			f.m[p] = append(f.m[p], id)
		}
	}
}

func (f *fabric) badlyClaimed() int {
	count := 0
	for _, ids := range f.m {
		if len(ids) > 1 {
			count++
		}
	}
	return count
}

func (f *fabric) safeOption() int {
	good := map[int]bool{}
	for _, id := range f.ids {
		good[id] = true
	}

	for _, ids := range f.m {
		if len(ids) <= 1 {
			continue
		}
		for _, id := range ids {
			delete(good, id)
		}
		if len(good) == 1 {
			break
		}
	}

	if len(good) != 1 {
		log.Fatalf("unexpected amount of goodness: %d", len(good))
	}

	// return only element in the set.
	var safeID int
	for id := range good {
		safeID = id
	}
	return safeID
}

func (f *fabric) Print() {
	var maxX, maxY int
	for p := range f.m {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	for x := 0; x <= maxY; x++ {
		for y := 0; y <= maxY; y++ {
			fmt.Print(len(f.m[xy{x, y}]))
		}
		fmt.Println()
	}
}
