package main

import (
	"fmt"
	"unsafe"
)

// Example of usage of merge function
// see utils.go and tests for more details
func main() {

	//check real size of array
	s := make([]int, 10000)
	fmt.Println("Size of []int", unsafe.Sizeof(s))

	defer elapsedTime("merge")()
	fmt.Println("Example: Merge of intervals")
	fmt.Printf("Result %s\n", mergeItervals([]Interval{Interval{25, 30}, Interval{2, 19}, Interval{14, 23}, Interval{8, 10}, Interval{4, 8}}))
}
