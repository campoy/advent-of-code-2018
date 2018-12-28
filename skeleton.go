package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

type data struct{}

func readInput(path string) (*data, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s: %v", path, err)
	}
	defer f.Close()

	var d data

	s := bufio.NewScanner(f)
	for s.Scan() {
		// do stuff
	}
	return &d, s.Err()
}
