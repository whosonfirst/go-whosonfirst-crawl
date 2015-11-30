package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-crawl"
	"runtime"
)

func main() {

     	var procs = flag.Int("processes", runtime.NumCPU() * 2, "The number of concurrent processes to use")
	
	flag.Parse()
	args := flag.Args()

	runtime.GOMAXPROCS(*procs)

	root := args[0]
	fmt.Println("crawl ", root)

	callback := func(path string, b crawl.JSONBlob) error {
		fmt.Println("inflated ", path)
		return nil
	}

	c := crawl.NewCrawler(root)
	_ = c.CrawlAndInflate(callback)
}
