package crawl

import (
	"fmt"
)

type CrawlError struct {
	Path    string
	Details error
}

func (e *CrawlError) Error() string {
	return e.String()
}

func (e *CrawlError) String() string {
	return fmt.Sprintf("Failed crawl for %s: %v", e.Path, e.Details)
}

type CallbackError struct {
	Path    string
	Details error
}

func (e *CallbackError) Error() string {
	return e.String()
}

func (e *CallbackError) String() string {
	return fmt.Sprintf("Failed crawl callback for %s: %v", e.Path, e.Details)
}

func NewCrawlError(path string, details error) *CrawlError {

	err := CrawlError{
		Path:    path,
		Details: details,
	}

	return &err
}

func NewCallbackError(path string, details error) *CallbackError {

	err := CallbackError{
		Path:    path,
		Details: details,
	}

	return &err
}
