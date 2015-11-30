prep:
	if test -d pkg; then rm -rf pkg; fi

self:	prep
	if test -d src/github.com/whosonfirst/go-whosonfirst-crawl; then rm -rf src/github.com/whosonfirst/go-whosonfirst-crawl; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-crawl
	cp crawl.go src/github.com/whosonfirst/go-whosonfirst-crawl/crawl.go

deps:   self
	go get -u "github.com/whosonfirst/walk"

fmt:
	go fmt cmd/*.go
	go fmt *.go

count:  self
	go build -o bin/wof-count cmd/wof-count.go

crawl:  self
	go build -o bin/wof-crawl cmd/wof-crawl.go

bin: crawl count
