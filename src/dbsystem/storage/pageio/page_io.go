package pageio

import (
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/lib/datastructs/appsync"
)

func NewPageIO(file dbfs.FileLike) (IO, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &pageIO{
		src:           file,
		lockMap:       appsync.NewResLockMap[page.Id](),
		pageCount:     uint32(fileInfo.Size() / pageSize),
		pageCountLock: appsync.SpinLock(0),
	}, err
}

var _ IO = &pageIO{}

type pageIO struct {
	src     dbfs.FileLike
	lockMap *appsync.ResLockMap[page.Id]

	pageCount     uint32
	pageCountLock appsync.SpinLock
}

func (p *pageIO) ReadPage(pageIndex page.Id, buffer []byte) error {
	if pageIndex >= p.pageCount {
		return NoPageErrorType
	}

	p.lockMap.RLockRes(pageIndex)
	defer p.lockMap.RUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.ReadAt(buffer, offset)

	if err != nil {
		return err
	}
	return nil
}

func (p *pageIO) FlushPage(pageIndex page.Id, pageData []byte) error {
	p.lockMap.WLockRes(pageIndex)
	defer p.lockMap.WUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.WriteAt(pageData, offset)
	if err != nil {
		return err
	}

	p.pageCountLock.Lock()
	if p.pageCount <= pageIndex {
		p.pageCount = pageIndex + 1
	}
	p.pageCountLock.Unlock()

	return err
}

func (p *pageIO) PageCount() uint32 {
	return p.pageCount
}

func (p *pageIO) PrepareNewPage() page.Id {
	p.pageCountLock.Lock()
	defer p.pageCountLock.Unlock()
	newPageId := p.pageCount
	p.pageCount++
	return newPageId
}

func (p *pageIO) Close() error {
	return p.src.Close()
}
