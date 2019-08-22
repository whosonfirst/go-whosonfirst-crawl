package crawl

import (
	walk "github.com/whosonfirst/walk"
	"os"
)

type CrawlFunc func(path string, info os.FileInfo) error

type Crawler struct {
	Root             string
	CrawlDirectories bool
}

func NewCrawler(path string) *Crawler {
	return &Crawler{
		Root:             path,
		CrawlDirectories: false,
	}
}

func (c Crawler) Crawl(cb CrawlFunc) error {

	walker := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && !c.CrawlDirectories {
			return nil
		}

		return cb(path, info)
	}

	return walk.Walk(c.Root, walker)
}
