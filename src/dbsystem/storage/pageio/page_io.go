package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/storage/dbfs"
	"sync"
)

func NewPageIO(file dbfs.FileLike) PageIO {
	return &pageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[dbbs.PageId](),
		pageCount:   0,
		newPageLock: &sync.Mutex{},
	}
}

func LoadPageIO(file dbfs.FileLike) PageIO {
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err.Error())
	}
	return &pageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[dbbs.PageId](),
		pageCount:   uint32(fileInfo.Size() / int64(pageSize)),
		newPageLock: &sync.Mutex{},
	}
}

var _ PageIO = &pageIO{}

type pageIO struct {
	src     dbfs.FileLike
	lockMap *appsync.ResLockMap[dbbs.PageId]

	pageCount   uint32
	newPageLock sync.Locker
}

func (p *pageIO) ReadPage(pageIndex dbbs.PageId, buffer []byte) error {
	p.lockMap.RLockRes(pageIndex)
	defer p.lockMap.RUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.ReadAt(buffer, offset)

	return err
}

func (p *pageIO) FlushPage(pageIndex dbbs.PageId, pageData []byte) error {
	p.lockMap.WLockRes(pageIndex)
	defer p.lockMap.WUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.WriteAt(pageData, offset)
	return err
}

func (p *pageIO) NewPage(pageData []byte) (dbbs.PageId, error) {
	p.newPageLock.Lock()
	defer p.newPageLock.Unlock()

	_, err := p.src.Write(pageData)
	if err != nil {
		return 0, err
	}
	p.pageCount++
	return p.pageCount, nil
}

func (p *pageIO) PageCount() uint32 {
	return p.pageCount
}

func (p *pageIO) Close() error {
	return p.src.Close()
}
