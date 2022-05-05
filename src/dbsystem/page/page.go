package page

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/io/bparse"
	"encoding/binary"
	"errors"
	"fmt"
)

const Size uint16 = 8192

func CreateEmptyPage(tableDef table.Definition) Page {
	page := make([]byte, Size)
	page[page]
	//todo implement
	
	return Page{
		table: tableDef,
		page:  ,
	}
}

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
	if p.TupleCount() <= tIndex {
		panic(fmt.Sprintf("Page has %d tuples but was requestd tuple with id: %d",
			p.TupleCount(), tIndex))
	}
	tuplePtr := p.tuplePtr(tIndex)
	tupleEndPtr := p.tupleEnd(tIndex)

	return Tuple{
		data:  p.page[tuplePtr:tupleEndPtr],
		table: p.table,
	}
}

func (p Page) LastTuplePtr() InPagePointer {
	pointerToLastPointer := bparse.Parse.UInt2(
		p.page[poLastTuplePtrPtr : poLastTuplePtrPtr+InPagePointerSize])

	return bparse.Parse.UInt2(p.page[pointerToLastPointer:InPagePointerSize])
}

func (p Page) TupleCount() uint16 {
	lastPointer := p.LastTuplePtr()
	if lastPointer == 0 {
		return 0
	}

	firstPointer := poFirstTuplePtr

	return (lastPointer-InPagePointer(firstPointer))/2 + 1
}

func (p Page) FreeSpace() uint16 {
	lastTuplePtr := p.LastTuplePtr()
	lastTuplePtrPtr := binary.LittleEndian.Uint16(
		p.page[poLastTuplePtrPtr : poLastTuplePtrPtr+InPagePointerSize])

	return lastTuplePtr - (lastTuplePtrPtr + InPagePointerSize - 1)
}

func (p Page) InsertTuple(tuple []byte) error {
	tupleLen := uint16(len(tuple))

	if tupleLen+InPagePointerSize > p.FreeSpace() {
		return errors.New(fmt.Sprintf("cannot fit tuple of size: %d to page with free space: %d",
			tupleLen, p.FreeSpace()))
	}

	// copy tuple
	lastTuplePtr := p.LastTuplePtr()
	newTuplePtr := lastTuplePtr - tupleLen
	copy(p.page[newTuplePtr:lastTuplePtr], tuple)

	// create new tuple pointer
	tupleLenBytes := make([]byte, 0, 2)
	binary.LittleEndian.PutUint16(tupleLenBytes, tupleLen)

	lastTuplePtrPtr := p.lastTuplePtrPtr()
	newTuplePtrPtr := lastTuplePtrPtr + InPagePointerSize
	copy(p.page[newTuplePtrPtr:newTuplePtrPtr+InPagePointerSize], tupleLenBytes)

	// reassign last pointer pointer
	copy(p.page[poLastTuplePtrPtr:poLastTuplePtrPtr+InPagePointerSize],
		tupleLenBytes)

	return nil
}

func (p Page) tupleEnd(index TupleIndex) InPagePointer {
	tuplePtr := p.tuplePtr(index)
	if tuplePtr == p.LastTuplePtr() {
		return Size
	}

	prevTuplePtr := p.tuplePtr(index - 1)
	return prevTuplePtr
}

func (p Page) tuplePtr(index TupleIndex) InPagePointer {
	ptrStart := poFirstTuplePtr + InPagePointerSize*index
	return bparse.Parse.UInt2(p.page[ptrStart : ptrStart+InPagePointerSize])
}

func (p Page) lastTuplePtrPtr() InPagePointer {
	return binary.LittleEndian.Uint16(
		p.page[poLastTuplePtr : poLastTuplePtr+InPagePointerSize])
}

func (p Page) updateHash() {
	//TODO implement me
	panic("not yet implemented")
}

type InPagePointer = uint16

const InPagePointerSize = 2

//page offsets
const (
	poPageHash        = 0                     // offset to page hash
	poLastTuplePtrPtr = poPageHash + 8        // offset to pointer pointing to last tuple pointer
	poLastTuplePtr    = poLastTuplePtrPtr + 2 // offset to pointer pointing to last tuple start
	poFirstTuplePtr   = poLastTuplePtr + 2    // offset to first of many tuple pointers
)
