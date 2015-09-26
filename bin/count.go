package main

import (
	"com.mapzen/whosonfirst"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("crawl ", root)

	var counter int64

	callback := func(path string, info os.FileInfo) error {

		if info.IsDir() {
			return nil
		}

		counter++
		return nil
	}

	t0 := time.Now()

	c := whosonfirst.NewCrawler(root)
	_ = c.Crawl(callback)

	t1 := float64(time.Since(t0)) / 1e9
	fmt.Printf("walked %d files in %.3f seconds\n", counter, t1)
}
