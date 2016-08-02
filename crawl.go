package crawl

import (
       "errors"
	"fmt"
	walk "github.com/whosonfirst/walk"
	"os"
)

type CrawlFunc func(path string, info os.FileInfo) error

type Crawler struct {
	Root string
	NFS bool
}

func NewCrawler(path string) *Crawler {
	return &Crawler{
		Root: path,
		NFS: false,
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

	var err errors.Error

	if c.NFS {
		err = walk.WalkNFS(c.Root, walker)
	} else {
		err = walk.Walk(c.Root, walker)
	}

	if err != nil {
		fmt.Printf("error: %s\n", err)
	}

	return nil
}
