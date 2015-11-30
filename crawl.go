package crawl

import (
	"encoding/json"
	"fmt"
	walk "github.com/whosonfirst/walk"
	"io/ioutil"
	"os"
)

type JSONBlob interface{}

type CrawlFunc func(path string, info os.FileInfo) error

type CrawlInflateFunc func(path string, blob JSONBlob) error

type Crawler struct {
	Root string
}

func NewCrawler(path string) *Crawler {
	return &Crawler{
		Root: path,
	}
}

func (c Crawler) Crawl(cb CrawlFunc) error {

	walker := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// ensure dot-geojson here

		_ = cb(path, info)
		return nil
	}

	err := walk.Walk(c.Root, walker)

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	return nil
}

func (c Crawler) CrawlAndInflate(cb CrawlInflateFunc) error {

	crawl_func := func(path string, info os.FileInfo) error {

		body, err := ioutil.ReadFile(path)

		if err != nil {
			fmt.Printf("failed to read %s, because %s", path, err)
			return err
		}

		var b JSONBlob
		err = json.Unmarshal(body, &b)

		if err != nil {
			fmt.Printf("failed to parse %s, because %s", path, err)
			return err
		}

		_ = cb(path, b)
		return nil
	}

	_ = c.Crawl(crawl_func)
	return nil
}
