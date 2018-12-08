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

	var twos, threes int
	s := bufio.NewScanner(f)
	for s.Scan() {
		counts := map[rune]int{}
		for _, r := range s.Text() {
			counts[r]++
		}

		var gotTwo, gotThree bool
		for _, c := range counts {
			if c == 2 {
				gotTwo = true
			}
			if c == 3 {
				gotThree = true
			}
		}

		if gotTwo {
			twos++
		}
		if gotThree {
			threes++
		}
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(twos * threes)
}
