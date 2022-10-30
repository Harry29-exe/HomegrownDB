package pageio

import (
	"HomegrownDB/dbsystem/storage/page"
	"io"
)

type IO interface {
	// ReadPage reads page with given index to provided buffer
	ReadPage(pageIndex page.Id, buffer []byte) error
	// FlushPage overrides pages at given page index with data from provided buffer
	FlushPage(pageIndex page.Id, pageData []byte) error
	// PageCount returns number of pages saved to disc
	PageCount() uint32

	io.Closer
}

type ResourceLockIO interface {
	RPage(pageId page.Id, buffer []byte) error
	ReleaseRPage(pageId page.Id)

	WPage(pageId page.Id, buffer []byte) error
	ReleaseWPage(pageId page.Id)

	Flush(pageId page.Id, pageData []byte) error
	NewPage(pageData []byte) (page.Id, error)

	PageCount() uint32
	io.Closer
}

var (
	pageSize = int64(page.Size)
)

var NoPageErrorType error = noPageError{}

type noPageError struct {
	pageId page.Id
}

func (n noPageError) Error() string {
	return "No page with given index"
}
