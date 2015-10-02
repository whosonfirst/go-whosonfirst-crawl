path:
	export GOPATH=`pwd`

count:  path
	go build -o bin/wof-count bin/wof-count.go

crawl:  path
	go build -o bin/wof-crawl bin/wof-crawl.go
