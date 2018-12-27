package aoc

func Sum(vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}
