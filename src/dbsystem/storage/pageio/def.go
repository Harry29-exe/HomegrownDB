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
	// PrepareNewPage creates space for future new page and returns id of future page
	PrepareNewPage() page.Id

	io.Closer
}

var (
	pageSize = int64(page.Size)
)

var NoPageErrorType error = noPageError{}

type noPageError struct{}

func (n noPageError) Error() string {
	return "No page with given index"
}
