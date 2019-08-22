package crawl

import (
	walk "github.com/whosonfirst/walk"
	"log"
	"os"
)

type CrawlFunc func(path string, info os.FileInfo) error

type Crawler struct {
	Root             string
	CrawlDirectories bool
	Strict bool
}

func NewCrawler(path string) *Crawler {
	return &Crawler{
		Root:             path,
		CrawlDirectories: false,
		Strict: true,
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

		err = cb(path, info)

		if c.Strict {
			return err
		}

		log.Printf("Callback failed for path '%s': %v\n", path, err)
		return nil
	}

	return walk.Walk(c.Root, walker)
}
