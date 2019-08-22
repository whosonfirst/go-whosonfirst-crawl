package crawl

import (
	"context"
	walk "github.com/whosonfirst/walk"
	_ "log"
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

	ctx := context.Background()

	return c.CrawlWithContext(ctx, cb)
}

func (c Crawler) CrawlWithContext(ctx context.Context, cb CrawlFunc) error {

	// this bit is important - see abouts about ctx.Done() and DoneError()
	// below in CrawlWithChannels (20190822/thisisaaronland)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	processing_ch := make(chan string)
	error_ch := make(chan error)
	done_ch := make(chan bool)

	go c.CrawlWithChannels(ctx, cb, processing_ch, error_ch, done_ch)

	for {
		select {
		case <-done_ch:
			return nil
		case <-processing_ch:
			// pass
		case err := <-error_ch:
			return err
		}
	}
}

func (c Crawler) CrawlWithChannels(ctx context.Context, cb CrawlFunc, processing_ch chan string, error_ch chan error, done_ch chan bool) {

	defer func() {
		done_ch <- true
	}()

	// note the bit with the DoneError() - if the `context.Context` object has signaled
	// that we're done we want to stop processing files but the only way to do that is
	// to send the `walk.Walk` object an error. In this case it's a special "done" error
	// that is not bubbled back up the stack to the caller (20190822/thisisaaronland)

	walker := func(path string, info os.FileInfo, err error) error {

		select {
		case <-ctx.Done():
			return NewDoneError()
		default:
			// pass
		}

		if err != nil {
			error_ch <- NewWalkError(path, err)
			return nil
		}

		if info.IsDir() && !c.CrawlDirectories {
			return nil
		}

		processing_ch <- path

		err = cb(path, info)

		if err != nil {
			error_ch <- NewCallbackError(path, err)
			return nil
		}

		return nil
	}

	err := walk.Walk(c.Root, walker)

	if err != nil && !IsDoneError(err) {
		error_ch <- NewCrawlError(c.Root, err)
	}
}
