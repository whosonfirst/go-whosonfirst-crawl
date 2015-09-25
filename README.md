# go-mapzen-whosonfirst-crawl

Experimental Go tools and libraries for crawling a directory of Who's On First data

## Usage

## Example

```
package main

import (
	"com.mapzen/whosonfirst"
	"flag"
	"fmt"
)

func main() {

	flag.Parse()
	args := flag.Args()

	root := args[0]
	fmt.Println("crawl ", root)

	callback := func(path string, b whosonfirst.JSONBlob) error {
		fmt.Println("inflated ", path)
		return nil
	}

	c := whosonfirst.NewCrawler(root)
	_ = c.CrawlAndInflate(callback)
}
```

## See also

* https://github.com/MichaelTJones/walk