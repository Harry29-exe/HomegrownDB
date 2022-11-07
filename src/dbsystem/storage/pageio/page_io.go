package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"errors"
)

func NewPageIO(file dbfs.FileLike) (IO, error) {
	if stat, err := file.Stat(); err != nil {
		return nil, err
	} else if stat.Size() != 0 {
		return nil, errors.New("to create new PageIO file must be empty")
	}

	return &pageIO{
		src:           file,
		lockMap:       appsync.NewResLockMap[page.Id](),
		pageCount:     0,
		pageCountLock: appsync.SpinLock(0),
	}, nil
}

func LoadPageIO(file dbfs.FileLike) (IO, error) {
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

func (p *pageIO) Close() error {
	return p.src.Close()
}
