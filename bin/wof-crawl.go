package main

import (
	"flag"
	"fmt"
	crawl "github.com/whosonfirst/go-whosonfirst-crawl/whosonfirst"
)

func main() {

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("crawl ", root)

	callback := func(path string, b crawl.JSONBlob) error {
		fmt.Println("inflated ", path)
		return nil
	}

	c := crawl.NewCrawler(root)
	_ = c.CrawlAndInflate(callback)
}
