package testmocks

import (
	"HomegrownDB/dbsystem/bdata"
	"HomegrownDB/dbsystem/io"
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

var pageSize = io.PageSize

func (t *TestInMemoryTableIO) ReadPage(pageIndex bdata.PageId, buffer []byte) error {
	pageStart := pageIndex * pageSize
	if int(pageStart) > len(t.pages) {
		return fmt.Errorf("no page with index: %d", pageIndex)
	}
	copy(buffer, t.pages[pageStart:pageStart+pageSize])

	return nil
}

func (t *TestInMemoryTableIO) FlushPage(pageIndex bdata.PageId, pageData []byte) error {
	pageStart := pageIndex * pageSize
	if int(pageStart) > len(t.pages) {
		return fmt.Errorf("no page with index: %d", pageIndex)
	}
	copy(t.pages[pageStart:pageStart+pageSize], pageData)

	return nil
}

func (t *TestInMemoryTableIO) NewPage(pageData []byte) (bdata.PageId, error) {
	t.pages = append(t.pages, pageData...)

	return bdata.PageId(len(t.pages))/pageSize - 1, nil
}

func (t *TestInMemoryTableIO) ReadBgPage(pageIndex bdata.PageId, buffer []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) FlushBgPage(pageIndex bdata.PageId, pageData []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) NewBgPage(pageData []byte) (bdata.PageId, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) ReadToastPage(pageIndex bdata.PageId, buffer []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) FlushToastPage(pageIndex bdata.PageId, pageData []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t *TestInMemoryTableIO) NewToastPage(pageData []byte) (bdata.PageId, error) {
	//TODO implement me
	panic("implement me")
}
