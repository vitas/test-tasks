package main

import (
	"fmt"
)

// Example of usage of merge function
// see utils.go for more details
func main() {
	defer elapsedTime("merge")()
	fmt.Println("Merge of intervals")
	fmt.Println(mergeItervals([]Interval{Interval{1, 15}, Interval{17, 20}, Interval{18, 20}}))
}
