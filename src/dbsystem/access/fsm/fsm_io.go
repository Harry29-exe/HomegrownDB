package fsm

import (
	"HomegrownDB/common/datastructs/appsync"
	_ "HomegrownDB/dbsystem/access/dbfs"
	"HomegrownDB/dbsystem/dbbs"
	_ "HomegrownDB/dbsystem/schema/table"
	"fmt"
	"os"
	"sync"
)

func loadIO(file *os.File) (*io, error) {
	return &io{
		file:        file,
		pageLock:    appsync.NewResLockMap[dbbs.PageId](),
		newPageLock: &sync.Mutex{},
	}, nil
}

func createNewIO(filepath string) (*io, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	pagesToCreate := 1
	for j := uint16(0); j < pageLayers-2; j++ {
		pagesToCreate *= int(layersInPage)
	}
	pagesToCreate += 1

	buff := make([]byte, pageSize)
	for j := 0; j < pagesToCreate; j++ {
		n, err := file.Write(buff)
		if err != nil {
			return nil, deleteOpenFileAfterErr(file, err)
		} else if n != int(pageSize) {
			panic(fmt.Sprintf(
				"expected that %d bytes will be written to file but %d bytes were written",
				pageSize, n),
			)
		}
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return loadIO(file)
}

type io struct {
	file        *os.File
	pageLock    *appsync.ResLockMap[dbbs.PageId]
	newPageLock *sync.Mutex
}

func (i *io) createInitialPages() error {
	i.file.Name()

	pagesToSave := 1
	for j := uint16(0); j < layersInPage-2; j++ {
		pagesToSave *= int(layersInPage)
	}
	pagesToSave += 1

	buff := make([]byte, pageSize)
	for j := 0; j < pagesToSave; j++ {
		_, err := i.file.Write(buff)
		if err != nil {
			return deleteOpenFileAfterErr(i.file, err)
		}
	}

	return nil
}

func (i *io) createNewPage() error {
	i.newPageLock.Lock()
	defer i.newPageLock.Unlock()

	writtenBytes, err := i.file.Write(make([]byte, pageSize))
	if err != nil {
		return err
	}
	if writtenBytes != int(pageSize) {
		return fmt.Errorf("bytes written to file(%d) are not equal to page size(%d)",
			writtenBytes, pageSize)
	}

	err = i.file.Close()
	if err != nil {
		return err
	}

	return nil
}

// flushPage saves existing page back to disc
// it assumes that page was acquired with WLock
func (i *io) flushPage(pageId dbbs.PageId, bytes []byte) error {
	writtenBytes, err := i.file.WriteAt(bytes, int64(pageId*uint32(pageSize)))
	if err != nil {
		return err
	}
	if writtenBytes != int(pageSize) {
		return fmt.Errorf("bytes written to file(%d) are not equal to page size(%d)",
			writtenBytes, pageSize)
	}

	return nil
}

func (i *io) releaseRPage(id dbbs.PageId) {
	i.pageLock.RUnlockRes(id)
}

func (i *io) releaseWPage(id dbbs.PageId) {
	i.pageLock.WUnlockRes(id)
}

func (i *io) getRPage(pageId dbbs.PageId, buffer []byte) ([]byte, error) {
	i.pageLock.RLockRes(pageId)

	return i.readPage(pageId, buffer)
}

func (i *io) getWPage(pageId dbbs.PageId, buffer []byte) ([]byte, error) {
	i.pageLock.WLockRes(pageId)

	return i.readPage(pageId, buffer)
}

func (i *io) readPage(pageId dbbs.PageId, buffer []byte) ([]byte, error) {
	n, err := i.file.ReadAt(buffer, int64(pageId)*int64(pageSize))
	if err != nil {
		return nil, err
	} else if n != int(pageSize) {
		return nil, fmt.Errorf(
			"read bytes(%d) are different that page size(%d)",
			n, pageSize,
		)
	}
	return buffer, nil
}

func deleteOpenFileAfterErr(file *os.File, causeErr error) error {
	if file == nil {
		return fmt.Errorf("provided file is nil, after error: %s", causeErr.Error())
	}

	fName := file.Name()
	err := file.Close()
	if err != nil {
		return fmt.Errorf("not expected error: %s\n "+
			"while closing file after error: %s",
			err.Error(),
			causeErr.Error(),
		)
	}

	err = os.Remove(fName)
	if err != nil {
		return fmt.Errorf("not expected error: %s\n "+
			"while deleting file after error: %s",
			err.Error(),
			causeErr.Error(),
		)
	}

	return causeErr
}
