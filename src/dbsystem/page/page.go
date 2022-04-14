package page

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/io/bparse"
)

const Size uint16 = 8192

type Page struct {
	table table.Definition
	page  []byte
}

type Id = uint32
type Tag struct {
	PageId  Id
	TableId table.Id
}

func (p Page) Tuple(tIndex TupleIndex) Tuple {
	tuplePointerLocation := pageFirstTuplePtr + InPagePointerSize* tIndex
	var tupleByteAfter InPagePointer
	if p.LastTuplePointer() == tuplePointerLocation {
		tupleByteAfter = Size
	} else {
		tupleByteAfter =
	}

	return Tuple{
		data:  p.page[tuplePointerLocation[]],
		table: nil,
	}
}

func (p Page) LastTuplePointer() InPagePointer {
	pointerToLastPointer := bparse.Parse.UInt2(
		p.page[pageLastTuplePtrPtr:])

	return bparse.Parse.UInt2(p.page[pointerToLastPointer:])
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

func (p Page) tupleStart(index TupleIndex) InPagePointer {
	return bparse.Parse.UInt2(p.page[
		pageFirstTuplePtr + InPagePointerSize * index:
		])
}

type InPagePointer = uint16

const InPagePointerSize = 2

const (
	pagePageHash        = 0                       //offset to page hash
	pageLastTuplePtrPtr = pagePageHash + 8        // offset to pointer pointing to last tuple pointer
	pageLastTuplePtr    = pageLastTuplePtrPtr + 2 // offset to pointer pointing to last tuple start
	pageFirstTuplePtr   = pageLastTuplePtr + 2    // offset to first of many tuple pointers
)
