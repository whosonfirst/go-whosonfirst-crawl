package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-crawl"
	"log"
	"os"
	"sync/atomic"
	"time"
)

func do_crawl(root string) error {

	var files int64
	var dirs int64

	callback := func(path string, info os.FileInfo) error {

		if info.IsDir() {
			atomic.AddInt64(&dirs, 1)
			return nil
		}

		atomic.AddInt64(&files, 1)
		return nil
	}

	fmt.Println("count files and directories in ", root)

	t0 := time.Now()

	defer func() {
		t1 := float64(time.Since(t0)) / 1e9
		fmt.Printf("walked %d files (and %d dirs) in %.3f seconds\n", files, dirs, t1)
	}()

	c := crawl.NewCrawler(root)
	return c.Crawl(callback)
}

func main() {

	flag.Parse()

	for _, root := range flag.Args() {

		err := do_crawl(root)

		if err != nil {
			log.Fatal(err)
		}

	}

}
