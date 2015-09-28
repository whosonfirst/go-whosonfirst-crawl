path:
	export GOPATH=`pwd`

count:  path
	go build -o bin/count bin/count.go

crawl:  path
	go build -o bin/crawl bin/crawl.go
