package pageio

import "HomegrownDB/dbsystem/dbbs"

type PageIO interface {
	// ReadPage reads page with given index to provided buffer
	ReadPage(pageIndex dbbs.PageId, buffer []byte) error
	// FlushPage overrides pages at given page index with data from provided buffer
	FlushPage(pageIndex dbbs.PageId, pageData []byte) error
	// NewPage saves provided buffer as new page and returns newly created page index
	NewPage(pageData []byte) (dbbs.PageId, error)
	PageCount() uint32
}

type RWPageIO interface {
	RPage(pageId dbbs.PageId, buffer []byte) error
	ReleaseRPage(pageId dbbs.PageId)

	WPage(pageId dbbs.PageId, buffer []byte) error
	ReleaseWPage(pageId dbbs.PageId)

	Flush(pageId dbbs.PageId) error
	NewPage(pageData []byte) (dbbs.PageId, error)

	PageCount() uint32
}

var (
	pageSize = int64(dbbs.PageSize)
)
