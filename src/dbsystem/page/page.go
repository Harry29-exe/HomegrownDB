package page

import "HomegrownDB/dbsystem/schema/table"

const Size uint32 = 8192

type Page struct {
	table table.Definition
	page  []byte
}

type Id = uint32
type Tag struct {
	PageId  Id
	TableId table.Id
}

func (p Page) Tuple(tupleIndex uint16) Tuple {

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
