package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	serial, err := readNumberFromArgs()
	if err != nil {
		log.Fatal(err)
	}

	var maxX, maxY, maxP int
	for y := 1; y <= 300-3+1; y++ {
		for x := 1; x <= 300-3+1; x++ {
			p := powerAtSquare(x, y, serial)
			if p > maxP {
				maxX, maxY, maxP = x, y, p
			}
		}
	}
	fmt.Printf("%d,%d\n", maxX, maxY)
}

func readNumberFromArgs() (int, error) {
	if len(os.Args) != 2 {
		return -1, fmt.Errorf("expected one argument, got %d", len(os.Args)-1)
	}
	return strconv.Atoi(os.Args[1])
}

func powerAtSquare(x, y, serial int) int {
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += powerAt(x+i, y+j, serial)
		}
	}
	return sum
}

func powerAt(x, y, serial int) int {
	id := x + 10
	power := id * (id*y + serial)
	return (power/100)%10 - 5
}
