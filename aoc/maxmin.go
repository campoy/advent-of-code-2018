package aoc

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func lt(i, j int) bool            { return i < j }
func Min(vs ...int) int           { return top(vs, lt, MaxInt) }
func MinArg(vs ...int) (i, v int) { return topArg(vs, lt) }

func gt(i, j int) bool            { return i > j }
func Max(vs ...int) int           { return top(vs, gt, MinInt) }
func MaxArg(vs ...int) (i, v int) { return topArg(vs, gt) }

func top(vs []int, less func(int, int) bool, z int) int {
	m := z
	for _, v := range vs {
		if less(v, m) {
			m = v
		}
	}
	return m
}

func topArg(vs []int, less func(int, int) bool) (i, v int) {
	if len(vs) == 0 {
		return -1, 0
	}

	minI, minV := 0, vs[0]
	for i, v := range vs {
		if less(v, minV) {
			minV = v
			minI = i
		}
	}
	return minI, minV
}
