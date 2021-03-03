package main

import (
	"sort"
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
	//FIXME check max value for end number
	if len(i) == 2 && i[0] < i[1] {
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
// we consider sll intervals are valid values
func mergeItervals(intervals []Interval) []Interval {

	if len(intervals) == 0 {
		return []Interval{}
	}

	// TBD check if all intervals are valid? speed penalty..

	// initialize merged array of intervals for result
	mergedIntervals := make([]Interval, 0, len(intervals))

	// sort array of intervals based on start number
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].StartNum() < intervals[j].StartNum()
	})

	// start with first interval
	currInterval := intervals[0]

	//iterate through all intervals starting from second
	for i := 1; i < len(intervals); i++ {
		//if next interval starts is bigger end number of previos interval
		// than next is outside (no overlap) previos interval
		if intervals[i].StartNum() > currInterval.EndNum() {
			// include current in the list as there is no overlap
			mergedIntervals = append(mergedIntervals, currInterval)
			// move to check next interval on next iteration
			currInterval = intervals[i]
		} else if intervals[i].EndNum() > currInterval.EndNum() {
			// extend current interval with next as it has bigger range (has overlap)
			currInterval.ExtendTo(intervals[i].EndNum())

		}
	}
	// do not forget to append current
	mergedIntervals = append(mergedIntervals, currInterval)
	return mergedIntervals
}
