self:
        if test -d src/github.com/whosonfirst/go-whosonfirst-crawl; then rm -rf src/github.com/whosonfirst/go-whosonfirst-crawl; fi
        mkdir src/github.com/whosonfirst/go-whosonfirst-crawl
        cp -r whosonfirst src/github.com/whosonfirst/go-whosonfirst-crawl/whosonfirst

deps:   self
	go get -u "github.com/MichaelTJones/walk"

fmt:
	go fmt bin/wof-count.go
	go fmt bin/wof-crawl.go
	go fmt whosonfirst/crawl.go

count:  self
	go build -o bin/wof-count bin/wof-count.go

crawl:  self
	go build -o bin/wof-crawl bin/wof-crawl.go
