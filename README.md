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

## To do

* Documentation
* Proper error handling
* Actually filter on things ending in `.geojson`
* Make it work with [whosonfirst.WOFFeature](https://github.com/whosonfirst/go-mapzen-whosonfirst/blob/master/src/com.mapzen/whosonfirst/place.go) or whatever that ends up being

## See also

* https://github.com/MichaelTJones/walk