## Testing Task

Write a console application what merges overlapped intervals.
Input list of intervals, output array of merged overlapped intervals and intervals what could not be merged

Example:
Input: [25,30] [2,19] [14, 23] [4,8]  Output: [2,23] [25,30]


### Goals
Try to achieve best algorithmic efficiency (complexity, memory usage, estimated time)
Write tests and documentation, answer questions

### Limitations and Constraints
- interval has only two int numbers (n, j)
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

if make is not installed, just build the app

`go build -o merge_interval *.go`

run example

`./merge_interval`

### How to run tests
`make test`

without make
`go test`

### Solution evoluation
- What is execution time?
  - I think in worst case time complexity is O( n log n )
- How can robustness be ensured, especially with regard to very large inputs?
  - Knowing a memory usage of of intervals we can estimate max count of intervals we can process
  - Do not forget about CPU usage, use a profiler to get a sense of CPU amd memory consumtion, use GOGC param to increase memory before CG
  - Do performace testing and make a final estimation
  - Consider about parallel computing, like map reduce systems
- What is memory consumption?
  - For a 2 dimensional array of "int" a total size is about 24 + (24 + 4 * i) * j, where array with intervals int[i][j]. 
Basically in 64bit system, "int" is 4 bytes, int[] is 24 + 4 * i, where 24 is size of empty "int" array in go 


