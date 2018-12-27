package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const totalWidth = 300

func main() {
	serial, err := readNumberFromArgs()
	if err != nil {
		log.Fatal(err)
	}

	var maxX, maxY, maxL, maxP int
	for y := 1; y <= totalWidth; y++ {
		for x := 1; x <= totalWidth; x++ {
			p, l := powerAtSquare(x, y, serial)
			if p > maxP {
				maxX, maxY, maxL, maxP = x, y, l, p
			}
		}
	}
	fmt.Printf("%d,%d,%d\n", maxX, maxY, maxL)
}

func readNumberFromArgs() (int, error) {
	if len(os.Args) != 2 {
		return -1, fmt.Errorf("expected one argument, got %d", len(os.Args)-1)
	}
	return strconv.Atoi(os.Args[1])
}

func powerAtSquare(x, y, serial int) (p, l int) {
	sum := powerAt(x, y, serial)
	maxSum := sum
	maxL := 1

	for l := 2; l+x <= totalWidth+1 && l+y <= totalWidth+1; l++ {
		for i := 0; i <= l-1; i++ {
			sum += powerAt(x+i, y+l-1, serial)
		}
		for i := 0; i < l-1; i++ {
			sum += powerAt(x+l-1, y+i, serial)
		}
		if sum > maxSum {
			maxSum = sum
			maxL = l
		}
	}

	return maxSum, maxL
}

func powerAt(x, y, serial int) int {
	id := x + 10
	power := id * (id*y + serial)
	return (power/100)%10 - 5
}
