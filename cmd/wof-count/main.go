package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync/atomic"
	"time"

	crawl "github.com/whosonfirst/go-whosonfirst-crawl"
)

func main() {
	procs := flag.Int("processes", runtime.NumCPU()*2, "The number of concurrent processes to use")

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("count files and directories in ", root)

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

	t0 := time.Now()

	c := crawl.NewCrawler(root)
	c.CallbackConcurrency = uint(*procs)
	c.CrawlDirectories = true

	err := c.Crawl(callback)
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	t1 := float64(time.Since(t0)) / 1e9
	fmt.Printf("walked %d files (and %d dirs) in %.3f seconds\n", files, dirs, t1)
}
