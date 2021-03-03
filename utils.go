package main

import (
	"fmt"
	"time"
)

// print execution time in ms
// example of usage: defer elapsedTime("my func")()
func elapsedTime(tag string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		fmt.Printf("%s execution time : %dms \n", tag, elapsed.Nanoseconds()/1000)
	}
}
