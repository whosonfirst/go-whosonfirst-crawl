package main

import (
	"whosonfirst/crawl"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("count files and directories in ", root)

	var files int64
	var dirs int64

	callback := func(path string, info os.FileInfo) error {

		if info.IsDir() {
			dirs ++
			return nil
		}

		files++
		return nil
	}

	t0 := time.Now()

	c := whosonfirst.NewCrawler(root)
	_ = c.Crawl(callback)

	t1 := float64(time.Since(t0)) / 1e9
	fmt.Printf("walked %d files (and %d dirs) in %.3f seconds\n", files, dirs, t1)
}
