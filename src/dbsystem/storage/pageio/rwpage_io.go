package pageio

import (
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/dbsystem/storage/dbfs"
	"HomegrownDB/dbsystem/storage/page"
	"errors"
	"sync"
)

func NewRWPageIO(file dbfs.FileLike) (ResourceLockIO, error) {
	if stat, err := file.Stat(); err != nil {
		return nil, err
	} else if stat.Size() != 0 {
		return nil, errors.New("to create new PageIO file must be empty")
	}

	return &rwPageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[page.Id](),
		pageCount:   0,
		newPageLock: &sync.Mutex{},
	}, nil
}

func LoadRWPageIO(file dbfs.FileLike) (ResourceLockIO, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	return &rwPageIO{
		src:         file,
		lockMap:     appsync.NewResLockMap[page.Id](),
		pageCount:   uint32(fileInfo.Size() / pageSize),
		newPageLock: &sync.Mutex{},
	}, err
}

var _ ResourceLockIO = &rwPageIO{}

type rwPageIO struct {
	src     dbfs.FileLike
	lockMap *appsync.ResLockMap[page.Id]

	pageCount   uint32
	newPageLock sync.Locker
}

func (io *rwPageIO) RPage(pageId page.Id, buffer []byte) error {
	io.lockMap.RLockRes(pageId)

	_, err := io.src.ReadAt(buffer, io.calcOffset(pageId))
	if err != nil {
		io.lockMap.RUnlockRes(pageId)
		return err
	}
	return nil
}

func (io *rwPageIO) ReleaseRPage(pageId page.Id) {
	io.lockMap.RUnlockRes(pageId)
}

func (io *rwPageIO) WPage(pageId page.Id, buffer []byte) error {
	io.lockMap.WLockRes(pageId)

	_, err := io.src.ReadAt(buffer, io.calcOffset(pageId))
	if err != nil {
		io.lockMap.WUnlockRes(pageId)
		return err
	}
	return nil
}

func (io *rwPageIO) ReleaseWPage(pageId page.Id) {
	io.lockMap.WUnlockRes(pageId)
}

func (io *rwPageIO) Flush(pageId page.Id, pageData []byte) error {
	_, err := io.src.WriteAt(pageData, io.calcOffset(pageId))
	return err
}

func (io *rwPageIO) NewPage(pageData []byte) (page.Id, error) {
	io.newPageLock.Lock()
	defer io.newPageLock.Unlock()

	_, err := io.src.Write(pageData)
	if err != nil {
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

func (io *rwPageIO) calcOffset(pageId page.Id) int64 {
	return pageSize * int64(pageId)
}
