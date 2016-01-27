package crawl

import (
	"fmt"
	walk "github.com/whosonfirst/walk"
	"os"
)

type CrawlFunc func(path string, info os.FileInfo) error

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

		_ = cb(path, info)
		return nil
	}

	err := walk.Walk(c.Root, walker)

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	return nil
}