package main

import (
	"strconv"
	"strings"
)

type Interval []int

func (i Interval) Valid() bool {
	if len(i) == 2 {
		return true
	}
	return false
}
func (i Interval) StartNum() int {
	if i.Valid() {
		return i[0]
	}
	return 0
}

func (i Interval) EndNum() int {
	if i.Valid() {
		return i[1]
	}
	return 0
}

func (i Interval) ExtendTo(num int) bool {
	if i.Valid() {
		i[1] = num
		return true
	}
	return false
}

func (i Interval) String() string {
	if !i.Valid() {
		return "()"
	}
	b := make([]string, len(i))
	for i, v := range i {
		b[i] = strconv.Itoa(v)
	}
	return "(" + strings.Join(b, "-") + ")"
}

func mergeItervals(intervals []Interval) []Interval {

	mergedIntervals := []Interval{}

	return mergedIntervals
}
