package crawl

import (
	"context"
	"runtime"

	"github.com/karrick/godirwalk"
	"golang.org/x/sync/errgroup"
)

type CrawlFunc func(path string, isDirectory bool) error

type Crawler struct {
	Root                string
	CrawlDirectories    bool
	CallbackConcurrency uint
	LexigraphicalSort   bool
	FollowSymbolicLinks bool
}

func NewCrawler(path string) *Crawler {
	cbConcurrency := uint(runtime.NumCPU() * 2)

	return &Crawler{
		Root:                path,
		CrawlDirectories:    false,
		CallbackConcurrency: cbConcurrency,
		LexigraphicalSort:   false,
		FollowSymbolicLinks: false,
	}
}

type pathInfo struct {
	path        string
	isDirectory bool
}

// Crawl walks through everything in Root, calling the provided callback
// function for every file (and optionally, every directory) as it finds them.
func (c Crawler) Crawl(cb CrawlFunc) error {
	// Make this chan double the length of the concurrency, so there's a bit of
	// buffer before between pushing a new file onto the chan and a worker picking
	// it up.
	pathChan := make(chan *pathInfo, c.CallbackConcurrency*2)

	g, ctx := errgroup.WithContext(context.Background())

	var i uint
	for i = 0; i < c.CallbackConcurrency; i++ {
		startCallbackWorker(ctx, g, pathChan, cb)
	}

	g.Go(func() error {
		walker := func(path string, directoryEntry *godirwalk.Dirent) error {
			if directoryEntry.IsDir() && !c.CrawlDirectories {
				return nil
			}

			p := &pathInfo{path: path, isDirectory: directoryEntry.IsDir()}

			select {
			case pathChan <- p:
				return nil
			case <-ctx.Done():
				// Stop in case of an error in one of the workers
				return ctx.Err()
			}
		}

		opts := &godirwalk.Options{
			Callback:            walker,
			Unsorted:            !c.LexigraphicalSort,
			FollowSymbolicLinks: c.FollowSymbolicLinks,
		}

		err := godirwalk.Walk(c.Root, opts)
		if err != nil {
			return err
		}

		// We're done, so close this channel out, which signals to the workers there's
		// nothing else to do
		close(pathChan)

		return nil
	})

	return g.Wait()
}

func startCallbackWorker(ctx context.Context, g *errgroup.Group, pathChan chan *pathInfo, cb CrawlFunc) {
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()

			case p, more := <-pathChan:
				if !more {
					return nil
				}

				err := cb(p.path, p.isDirectory)
				if err != nil {
					return err
				}
			}
		}
	})
}
