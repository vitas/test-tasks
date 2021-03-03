package main

import (
	"strconv"
	"strings"
)

// define interval type, just for convience
type Interval []int

//compare two intervals
func (i Interval) Equals(ti Interval) bool {
	if len(i) == len(ti) && i.Valid() && ti.Valid() && i.StartNum() == ti.StartNum() && i.EndNum() == ti.EndNum() {
		return true
	}
	return false
}

// check if interval has valid , has start and end numbers
func (i Interval) Valid() bool {
	if len(i) == 2 {
		return true
	}
	return false
}

//get start number of interval
func (i Interval) StartNum() int {
	if i.Valid() {
		return i[0]
	}
	return 0
}

// get end number of interval
func (i Interval) EndNum() int {
	if i.Valid() {
		return i[1]
	}
	return 0
}

//extend end number of interval
func (i Interval) ExtendTo(num int) bool {
	if i.Valid() {
		i[1] = num
		return true
	}
	return false
}

// toString for values print
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

// merges overlapped intervals and keeps non overlapped intervals
func mergeItervals(intervals []Interval) []Interval {

	mergedIntervals := []Interval{}

	return mergedIntervals
}
