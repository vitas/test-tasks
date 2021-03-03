package main

import (
	"fmt"
	"testing"
)

// Test valid interval values
func TestValidIntervalPaar(t *testing.T) {
	validIntervals := []Interval{Interval{0, 20}, Interval{-5, 10}, Interval{5, 15}}
	for _, interval := range validIntervals {
		if !interval.Valid() {
			t.Errorf("Invalid interval value detected: %s", interval)
		}
	}
}

// Test invalid interval values
func TestInvalidIntervalPaar(t *testing.T) {
	invalidIntervals := []Interval{Interval{}, Interval{20, 0}, Interval{1}, Interval{1, 2, 3}, Interval{1, 1}}
	for _, interval := range invalidIntervals {
		if interval.Valid() {
			t.Errorf("Invalid interval value not detected: %s", interval)
		}
	}
}

// Tests if input parameter is empty interval array
func TestEmptyInput(t *testing.T) {
	merged := mergeItervals([]Interval{})
	len := len(merged)
	if len != 0 {
		t.Errorf("Invalid input interval array, got: %d, want: %d", len, 0)
	}
}

func TestOnlyOneInteval(t *testing.T) {
	defer elapsedTime("TestOneInterval")()
	inputInterval := []Interval{Interval{3, 45}}
	merged := mergeItervals(inputInterval)
	len := len(merged)
	if len != 1 {
		t.Errorf("Invalid input interval array, got: %d, want: %d", len, 1)
	}
}

//test merge of intervals
func TestIntevals(t *testing.T) {
	defer elapsedTime("TestIntervals")()

	testInputInterval := []Interval{Interval{25, 30}, Interval{2, 19}, Interval{14, 23}, Interval{8, 10}, Interval{4, 8}}
	testOutputInterval := []Interval{Interval{2, 23}, Interval{25, 30}}

	// some console debug info if fails
	fmt.Printf("input intervals: %s\n", testInputInterval)
	merged := mergeItervals(testInputInterval)
	fmt.Printf("merged intervals: %s\n", merged)
	fmt.Printf("expected output intervals: %s\n", testOutputInterval)

	lenMerged := len(merged)
	lenExpected := len(testOutputInterval)
	if lenMerged != lenExpected {
		t.Errorf("Invalid input interval array, got: %d, want: %d", lenMerged, lenExpected)
	} else {
		for i, expected := range testOutputInterval {
			if !expected.Equals(merged[i]) {
				t.Errorf("Invalid interval value detected: %s", merged[i])
			}
		}
	}
}
