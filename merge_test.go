package main

import "testing"

//TBD add range on vaslid values, inklusive negative ranges
// Test valid interval values
func TestIntervalPaar(t *testing.T) {
	interval := Interval{1, 10}
	if !interval.Valid() {
		t.Errorf("Invalid interval value detected")
	}
}

//TBD add range of invalid values
// Test invalid interval values
func TestIvalidIntervalPaar(t *testing.T) {
	interval := Interval{1, 10, 34}
	if interval.Valid() {
		t.Errorf("Invalid interval value not detected")
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
