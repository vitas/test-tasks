## Testing Task

Write a console application what merges overlapped intervals.
Input list of intervals, output list of merged overlapped intervals and intervals what could not be merged

Example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]


### Goals
Try to achieve best algorithmic efficiency (complexity, memory usage, estimated time)

### Limitations and Constraints
- interval has only two numbers (n, j)
- min two intervals could attempted to merge
- intervals array length should be between 1..n
- start of interval has to be <= end of interval
- end of interval <= n
- n ..how big is the value? 10000, more?

### Selected language
golang

### Make and build

build project
`make`

if make is not installed, just build
`go build -o merge_interval *.go`

run example

`./merge_interval`

### How to run tests
`make test`

without make
`go test`

### Solution evoluation
