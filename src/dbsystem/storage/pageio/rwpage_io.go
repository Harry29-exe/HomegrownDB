package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"sync"
)

var _ RWPageIO = &rwPageIO{}

type rwPageIO struct {
	src     dbfs.FileLike
	lockMap appsync.ResLockMap[dbbs.PageId]

	pageCount   uint32
	newPageLock sync.Locker
}

func (io *rwPageIO) RPage(pageId dbbs.PageId, buffer []byte) error {
	io.lockMap.RLockRes(pageId)

	_, err := io.src.ReadAt(buffer, io.calcOffset(pageId))
	if err != nil {
		io.lockMap.RUnlockRes(pageId)
		return err
	}
	return nil
}

func (io *rwPageIO) ReleaseRPage(pageId dbbs.PageId) {
	io.lockMap.RLockRes(pageId)
}

func (io *rwPageIO) WPage(pageId dbbs.PageId, buffer []byte) error {
	io.lockMap.WLockRes(pageId)

	_, err := io.src.ReadAt(buffer, io.calcOffset(pageId))
	if err != nil {
		io.lockMap.WUnlockRes(pageId)
		return err
	}
	return nil
}

func (io *rwPageIO) ReleaseWPage(pageId dbbs.PageId) {
	io.lockMap.WUnlockRes(pageId)
}

func (io *rwPageIO) Flush(pageId dbbs.PageId, pageData []byte) error {
	_, err := io.src.WriteAt(pageData, io.calcOffset(pageId))
	return err
}

func (io *rwPageIO) NewPage(pageData []byte) (dbbs.PageId, error) {
	io.newPageLock.Lock()
	_, err := io.src.Write(pageData)
	if err != nil {
		io.newPageLock.Unlock()
		return 0, err
	}
	io.pageCount++
	return io.pageCount - 1, nil
}

func (io *rwPageIO) PageCount() uint32 {
	return io.pageCount
}

func (io *rwPageIO) Close() error {
	return io.src.Close()
}

func (io *rwPageIO) calcOffset(pageId dbbs.PageId) int64 {
	return pageSize * int64(pageId)
}
