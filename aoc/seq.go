package aoc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
)

func ReadIntSeqFromFile(path string) (*IntSeq, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var nums []int
	for _, p := range bytes.Fields(bs) {
		n, err := strconv.Atoi(string(p))
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %v", p, err)
		}
		nums = append(nums, n)
	}
	return &IntSeq{nums}, nil
}

type IntSeq struct{ nums []int }

func NewIntSeq(nums ...int) *IntSeq { return &IntSeq{nums} }

func (s *IntSeq) Empty() bool {
	return len(s.nums) == 0
}

func (s *IntSeq) Next() int {
	r := s.nums[0]
	s.nums = s.nums[1:]
	return r
}
