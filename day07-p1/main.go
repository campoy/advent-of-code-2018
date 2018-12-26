package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	g, err := readGraph("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	h := g.heap()
	for {
		n := h.pop()
		if n == nil {
			fmt.Println()
			return
		}
		fmt.Print(n.name)
	}
}

type node struct {
	name    string
	deps    int
	follows []string
}

type graph map[string]*node

func readGraph(path string) (graph, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	g := make(graph)

	s := bufio.NewScanner(f)
	for s.Scan() {
		var a, b string
		fmt.Sscanf(s.Text(), "Step %s must be finished before step %s can begin.", &a, &b)
		g.addDep(a, b)
	}

	return g, s.Err()
}

func (g graph) addDep(a, b string) {
	g.addNode(a)
	g.addNode(b)
	g[b].deps++
	g[a].follows = append(g[a].follows, b)
}

func (g graph) addNode(name string) {
	if _, ok := g[name]; !ok {
		g[name] = &node{name: name}
	}
}

func (g graph) heap() heap {
	h := heap{g: g}
	for _, n := range g {
		h.nodes = append(h.nodes, n)
	}
	return h
}

type heap struct {
	g     graph
	nodes []*node
}

func (h *heap) pop() *node {
	if len(h.nodes) == 0 {
		return nil
	}

	sort.Slice(h.nodes, func(i, j int) bool {
		li, lj := h.nodes[i].deps, h.nodes[j].deps
		if li < lj {
			return true
		} else if li > lj {
			return false
		}
		return h.nodes[i].name < h.nodes[j].name
	})

	r := h.nodes[0]
	h.nodes = h.nodes[1:]

	// Update the number of dependencies.
	// The follows are not relevant.
	for _, f := range r.follows {
		h.g[f].deps--
	}

	return r
}
