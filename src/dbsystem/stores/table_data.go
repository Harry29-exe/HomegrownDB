package stores

import (
	"HomegrownDB/dbsystem"
	"HomegrownDB/dbsystem/bdata"
	"fmt"
	"os"
	"sync"
)

const PageSize = uint32(dbsystem.PageSize)

type TableDataStore interface {
	ReadPage(pageIndex bdata.PageId, buffer []byte) error
	FlushPage(pageIndex bdata.PageId, pageData []byte) error
	NewPage(pageData []byte) (bdata.PageId, error)

	ReadBgPage(pageIndex bdata.PageId, buffer []byte) error
	FlushBgPage(pageIndex bdata.PageId, pageData []byte) error
	NewBgPage(pageData []byte) (bdata.PageId, error)

	ReadToastPage(pageIndex bdata.PageId, buffer []byte) error
	FlushToastPage(pageIndex bdata.PageId, pageData []byte) error
	NewToastPage(pageData []byte) (bdata.PageId, error)
}

func SingleDiscTableDataStore(pathToTableDir string) (TableDataStore, error) {
	pagesFile, err := os.Open(pathToTableDir + "/pages.hdbd")
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

	return &tableDataStore{
		pagesFile:          pagesFile,
		pagesFileLock:      &sync.RWMutex{},
		bgPagesFile:        bgPagesFile,
		bgPagesFileLock:    &sync.RWMutex{},
		toastPagesFile:     toastPagesFile,
		toastPagesFileLock: &sync.RWMutex{},
	}, nil
}

type tableDataStore struct {
	pagesFile     *os.File
	pagesFileLock *sync.RWMutex

	bgPagesFile     *os.File
	bgPagesFileLock *sync.RWMutex

	toastPagesFile     *os.File
	toastPagesFileLock *sync.RWMutex
}

func (t *tableDataStore) ReadPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.pagesFile, t.pagesFileLock)
}

func (t *tableDataStore) FlushPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.pagesFile, t.pagesFileLock)
}

func (t *tableDataStore) NewPage(pageData []byte) (bdata.PageId, error) {
	return newPage(pageData, t.pagesFile, t.pagesFileLock)
}

func (t *tableDataStore) ReadBgPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.bgPagesFile, t.bgPagesFileLock)
}

func (t *tableDataStore) FlushBgPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.bgPagesFile, t.bgPagesFileLock)
}

func (t *tableDataStore) NewBgPage(pageData []byte) (bdata.PageId, error) {
	return newPage(pageData, t.bgPagesFile, t.bgPagesFileLock)
}

func (t *tableDataStore) ReadToastPage(pageIndex bdata.PageId, buffer []byte) error {
	return readPage(pageIndex, buffer, t.toastPagesFile, t.toastPagesFileLock)
}

func (t *tableDataStore) FlushToastPage(pageIndex bdata.PageId, pageData []byte) error {
	return flushPage(pageIndex, pageData, t.toastPagesFile, t.toastPagesFileLock)
}

func (t *tableDataStore) NewToastPage(pageData []byte) (bdata.PageId, error) {
	return newPage(pageData, t.toastPagesFile, t.toastPagesFileLock)
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
