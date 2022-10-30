package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"errors"
	"sync"
)

func NewPageIO(file dbfs.FileLike) (IO, error) {
	if stat, err := file.Stat(); err != nil {
		return nil, err
	} else if stat.Size() != 0 {
		return nil, errors.New("to create new PageIO file must be empty")
	}

	return &pageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[page.Id](),
		pageCount:   0,
		newPageLock: &sync.Mutex{},
	}, nil
}

func LoadPageIO(file dbfs.FileLike) (IO, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &pageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[page.Id](),
		pageCount:   uint32(fileInfo.Size() / pageSize),
		newPageLock: &sync.Mutex{},
	}, err
}

var _ IO = &pageIO{}

type pageIO struct {
	src     dbfs.FileLike
	lockMap *appsync.ResLockMap[page.Id]

	pageCount   uint32
	newPageLock sync.Locker
}

func (p *pageIO) ReadPage(pageIndex page.Id, buffer []byte) error {
	p.lockMap.RLockRes(pageIndex)
	defer p.lockMap.RUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.ReadAt(buffer, offset)

	return err
}

func (p *pageIO) FlushPage(pageIndex page.Id, pageData []byte) error {
	p.lockMap.WLockRes(pageIndex)
	defer p.lockMap.WUnlockRes(pageIndex)

	offset := int64(pageIndex) * pageSize
	_, err := p.src.WriteAt(pageData, offset)
	return err
}

func (p *pageIO) PageCount() uint32 {
	return p.pageCount
}

func (p *pageIO) Close() error {
	return p.src.Close()
}
