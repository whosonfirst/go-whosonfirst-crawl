# go-mapzen-whosonfirst-crawl

Go tools and libraries for crawling a directory of Who's On First data

## Usage

## Example

```
package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-crawl"
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

	c := whosonfirst.NewCrawler(root)
	_ = c.CrawlAndInflate(callback)
}
```

## To do

* Documentation
* Proper error handling
* Remove GeoJSON specific stuff (or at least move it in to its own little playground)

## See also

* https://github.com/MichaelTJones/walk