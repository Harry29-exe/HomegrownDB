package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"sync"
)

var _ PageIO = &pageIO{}

type pageIO struct {
	src     dbfs.FileLike
	lockMap appsync.ResLockMap[dbbs.PageId]

	pageCount   uint32
	newPageLock sync.Locker
}

func (p pageIO) ReadPage(pageIndex dbbs.PageId, buffer []byte) error {
	p.lockMap.RLockRes(pageIndex)
	defer p.lockMap.RUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.ReadAt(buffer, offset)

	return err
}

func (p pageIO) FlushPage(pageIndex dbbs.PageId, pageData []byte) error {
	p.lockMap.WLockRes(pageIndex)
	defer p.lockMap.WUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.WriteAt(pageData, offset)
	return err
}

func (p pageIO) NewPage(pageData []byte) (dbbs.PageId, error) {
	p.newPageLock.Lock()
	defer p.newPageLock.Unlock()

	_, err := p.src.Write(pageData)
	if err != nil {
		return 0, err
	}
	p.pageCount++
	return p.pageCount, nil
}

func (p pageIO) PageCount() uint32 {
	return p.pageCount
}
