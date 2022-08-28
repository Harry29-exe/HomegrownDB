package access

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/bdata"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

const PageSize = uint32(dbsystem.PageSize)

type TableDataIO interface {
	// ReadPage reads page with given index to provided buffer
	ReadPage(pageIndex bdata.PageId, buffer []byte) error
	// FlushPage overrides pages at given page index with data from provided buffer
	FlushPage(pageIndex bdata.PageId, pageData []byte) error
	// NewPage saves provided buffer as new page and returns newly created page index
	NewPage(pageData []byte) (bdata.PageId, error)
	PageCount() uint32

	ReadBgPage(pageIndex bdata.PageId, buffer []byte) error
	FlushBgPage(pageIndex bdata.PageId, pageData []byte) error
	NewBgPage(pageData []byte) (bdata.PageId, error)
	BgPageCount() uint32

	ReadToastPage(pageIndex bdata.PageId, buffer []byte) error
	FlushToastPage(pageIndex bdata.PageId, pageData []byte) error
	NewToastPage(pageData []byte) (bdata.PageId, error)
	ToastPageCount() uint32
}

func SingleDiscTableDataIO(pathToTableDir string) (TableDataIO, error) {
	pagesFile, err := os.Open(pathToTableDir + "/pages.hdbd") //todo read page count for all 3 page types
	if err != nil {
		return nil, err
	}
	bgPagesFile, err := os.Open(pathToTableDir + "/bg_pages.hdbd")
	if err != nil {
		return nil, err
	}
	toastPagesFile, err := os.Open(pathToTableDir + "/toast_pages.hdbd")
	if err != nil {
		return nil, err
	}

	return &tableDataIO{
		pagesFile:          pagesFile,
		pagesFileLock:      &sync.RWMutex{},
		bgPagesFile:        bgPagesFile,
		bgPagesFileLock:    &sync.RWMutex{},
		toastPagesFile:     toastPagesFile,
		toastPagesFileLock: &sync.RWMutex{},
	}, nil
}

type tableDataIO struct {
	pagesFile     *os.File
	pagesFileLock *sync.RWMutex
	pageCount     *uint32

	bgPagesFile     *os.File
	bgPagesFileLock *sync.RWMutex
	bgPageCount     *uint32

	toastPagesFile     *os.File
	toastPagesFileLock *sync.RWMutex
	toastPageCount     *uint32
}

func (t *tableDataIO) ReadPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.pagesFile, t.pagesFileLock)
}

func (t *tableDataIO) FlushPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.pagesFile, t.pagesFileLock)
}

func (t *tableDataIO) NewPage(pageData []byte) (bdata.PageId, error) {
	id, err := newPage(pageData, t.pagesFile, t.pagesFileLock)
	if err == nil {
		atomic.AddUint32(t.pageCount, 1)
	}
	return id, err
}

func (t *tableDataIO) PageCount() uint32 {
	return *t.pageCount
}

func (t *tableDataIO) ReadBgPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.bgPagesFile, t.bgPagesFileLock)
}

func (t *tableDataIO) FlushBgPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.bgPagesFile, t.bgPagesFileLock)
}

func (t *tableDataIO) NewBgPage(pageData []byte) (bdata.PageId, error) {
	id, err := newPage(pageData, t.bgPagesFile, t.bgPagesFileLock)
	if err == nil {
		atomic.AddUint32(t.bgPageCount, 1)
	}
	return id, err
}

func (t *tableDataIO) BgPageCount() uint32 {
	return *t.bgPageCount
}

func (t *tableDataIO) ReadToastPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.toastPagesFile, t.toastPagesFileLock)
}

func (t *tableDataIO) FlushToastPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.toastPagesFile, t.toastPagesFileLock)
}

func (t *tableDataIO) NewToastPage(pageData []byte) (bdata.PageId, error) {
	id, err := newPage(pageData, t.toastPagesFile, t.toastPagesFileLock)
	if err == nil {
		atomic.AddUint32(t.toastPageCount, 1)
	}
	return id, err
}

func (t *tableDataIO) ToastPageCount() uint32 {
	return *t.toastPageCount
}

func readPage(pageIndex bdata.PageId, buffer []byte, pagesFile *os.File, fileLock *sync.RWMutex) error {
	fileLock.RLock()
	defer fileLock.RUnlock()
	_, err := pagesFile.ReadAt(buffer, int64(PageSize*pageIndex))
	if err != nil {
		return err
	}

	return nil
}

func newPage(pageData []byte, pagesFile *os.File, fileLock *sync.RWMutex) (bdata.PageId, error) {
	fileLock.Lock()
	defer fileLock.Unlock()

	newPageId := getNumberOfPages(pagesFile) + 1
	n, err := pagesFile.Write(pageData)
	if err != nil && n == 0 {
		return 0, err
	} else if err != nil {
		fSizeBeforeWrite := int64(newPageId-1) * int64(PageSize)
		truncateErr := pagesFile.Truncate(fSizeBeforeWrite)
		if truncateErr != nil {
			panic(fmt.Sprintf("During writing to file: %s, ocured following error during write %s, after attempt to rollback changes but error: %s occured, this is critical error",
				pagesFile.Name(), err.Error(), truncateErr.Error()))
		}
		return 0, err
	}

	return newPageId, nil
}

func flushPage(pageIndex bdata.PageId, pageData []byte, pagesFile *os.File, fileLock *sync.RWMutex) error {
	fileLock.Lock()
	defer fileLock.Unlock()

	pageStart := int64(pageIndex) * int64(PageSize)
	_, err := pagesFile.WriteAt(pageData, pageStart)
	if err != nil {
		return err //todo better error handling
	}

	return nil
}

func getNumberOfPages(file *os.File) uint32 {
	info, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %s", err.Error()))
	}

	return uint32(info.Size() / int64(PageSize))
}
