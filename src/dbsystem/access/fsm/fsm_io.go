package fsm

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/math2"
	_ "HomegrownDB/dbsystem/access/dbfs"
	"HomegrownDB/dbsystem/dbbs"
	_ "HomegrownDB/dbsystem/schema/table"
	"fmt"
	"os"
	"sync"
)

func loadIO(file *os.File) (*io, error) {
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	return &io{
		file:        file,
		pageLock:    appsync.NewResLockMap[dbbs.PageId](),
		newPageLock: &sync.Mutex{},
		pages:       uint32(fileStat.Size() / int64(pageSize)),
	}, nil
}

func createNewIO(filepath string) (*io, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	pagesToCreate := 1
	for j := 1; j < pageLayers-1; j++ {
		pagesToCreate += math2.Power(int(leafNodeCount), j)
	}

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
	file     *os.File
	pageLock *appsync.ResLockMap[dbbs.PageId]

	pages       uint32
	newPageLock *sync.Mutex
}

func (i *io) createNewPage() {
	i.newPageLock.Lock()
	defer i.newPageLock.Unlock()

	writtenBytes, err := i.file.Write(make([]byte, pageSize))
	if err != nil {
		panic(err.Error())
	} else if writtenBytes != int(pageSize) {
		panic(fmt.Sprintf("bytes written to file(%d) are not equal to page size(%d)",
			writtenBytes, pageSize))
	}
}

// flushPage saves existing page back to disc
// it assumes that page was acquired with WLock
func (i *io) flushPage(pageId dbbs.PageId, bytes []byte) {
	writtenBytes, err := i.file.WriteAt(bytes, int64(pageId*uint32(pageSize)))
	if err != nil {
		panic(err.Error())
	} else if writtenBytes != int(pageSize) {
		panic(fmt.Sprintf("bytes written to file(%d) are not equal to page size(%d)",
			writtenBytes, pageSize))
	}
}

func (i *io) releaseRPage(id dbbs.PageId) {
	i.pageLock.RUnlockRes(id)
}

func (i *io) releaseWPage(id dbbs.PageId) {
	i.pageLock.WUnlockRes(id)
}

func (i *io) getRPage(pageId dbbs.PageId, buffer []byte) page {
	i.pageLock.RLockRes(pageId)

	err := i.readPage(pageId, buffer)
	if err != nil {
		i.pageLock.RUnlockRes(pageId)
		panic(err.Error())
	}
	return page{
		header: buffer[:headerSize],
		data:   buffer[headerSize:],
	}
}

func (i *io) getWPage(pageId dbbs.PageId, buffer []byte) page {
	i.pageLock.WLockRes(pageId)

	for pageId >= i.pages {
		err := i.createPage()
		if err != nil {
			panic(err.Error())
		}
	}

	err := i.readPage(pageId, buffer)
	if err != nil {
		i.pageLock.WUnlockRes(pageId)
		panic(err.Error())
	}
	return page{
		header: buffer[:headerSize],
		data:   buffer[headerSize:],
	}
}

func (i *io) createPage() error {
	i.newPageLock.Lock()
	defer i.newPageLock.Unlock()

	buff := make([]byte, pageSize)
	stat, err := i.file.Stat()
	if err != nil {
		return err
	}
	_, err = i.file.WriteAt(buff, stat.Size())
	if err != nil {
		return err
	}
	i.pages += 1
	return nil
}

func (i *io) readPage(pageId dbbs.PageId, buffer []byte) error {
	n, err := i.file.ReadAt(buffer, int64(pageId)*int64(pageSize))
	if err != nil {
		return err
	} else if n != int(pageSize) {
		return fmt.Errorf(
			"read bytes(%d) are different that page size(%d)",
			n, pageSize,
		)
	}
	return nil
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
