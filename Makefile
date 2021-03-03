build:
	go build -o merge_interval *.go

clean:
	rm merge_interval

test:
	go test
