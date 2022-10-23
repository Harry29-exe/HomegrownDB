package tstructs

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/storage/page"
	"fmt"
)

func NewInMemoryTableIO() *TestInMemoryTableIO {
	return &TestInMemoryTableIO{
		pages:      nil,
		bgPages:    nil,
		toastPages: nil,
	}
}

type TestInMemoryTableIO struct {
	pages      []byte
	bgPages    []byte
	toastPages []byte
}

var pageSize = access.PageSize

func (t *TestInMemoryTableIO) ReadPage(pageIndex page.Id, buffer []byte) error {
	pageStart := pageIndex * pageSize
	if int(pageStart) > len(t.pages) {
		return fmt.Errorf("no page with index: %d", pageIndex)
	}
	copy(buffer, t.pages[pageStart:pageStart+pageSize])

	return nil
}

func (t *TestInMemoryTableIO) FlushPage(pageIndex page.Id, pageData []byte) error {
	pageStart := pageIndex * pageSize
	if int(pageStart) > len(t.pages) {
		return fmt.Errorf("no page with index: %d", pageIndex)
	}
	copy(t.pages[pageStart:pageStart+pageSize], pageData)

	return nil
}

func (t *TestInMemoryTableIO) NewPage(pageData []byte) (page.Id, error) {
	t.pages = append(t.pages, pageData...)

	return page.Id(len(t.pages))/pageSize - 1, nil
}

func (t *TestInMemoryTableIO) PageCount() uint32 {
	return uint32(len(t.bgPages)) / pageSize
}

func (t *TestInMemoryTableIO) ReadBgPage(pageIndex page.Id, buffer []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) FlushBgPage(pageIndex page.Id, pageData []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) NewBgPage(pageData []byte) (page.Id, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) BgPageCount() uint32 {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) ReadToastPage(pageIndex page.Id, buffer []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) FlushToastPage(pageIndex page.Id, pageData []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) NewToastPage(pageData []byte) (page.Id, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) ToastPageCount() uint32 {
	//TODO implement me
	panic("implement me")
}
