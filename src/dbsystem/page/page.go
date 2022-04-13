package page

import "HomegrownDB/dbsystem/schema/table"

const Size uint32 = 8192

type Page struct {
	table table.Definition
	page  []byte
}

type Id struct {
	PageId      uint32
	LinePointer uint16
}

func (p Page) Tuple(tupleIndex uint16) Tuple {
	//TODO implement me
	panic("implement me")
}

func (p Page) TupleCount() uint16 {
	//TODO implement me
	panic("implement me")
}

func (p Page) FreeSpace() uint16 {
	//TODO implement me
	panic("implement me")
}

func (p Page) InsertTuple(data []byte) error {
	//TODO implement me
	panic("implement me")
}

type InPagePointer = uint16
