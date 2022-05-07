package page

// For better understanding of struct Page
// it's recommended to view page.doc.svg diagram
// which shows how page's binary representations looks like

import (
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/io/bparse"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
)

//todo add handling for inserting into empty page
const Size uint16 = 8192

func CreateEmptyPage(tableDef table.Definition) Page {
	rawPage := make([]byte, Size)
	uint16Zero := make([]byte, 2)
	binary.LittleEndian.PutUint16(uint16Zero, 0)

	copy(rawPage[poPrtToLastTuplePtr:poPrtToLastTuplePtr+InPagePointerSize], uint16Zero)
	copy(rawPage[poPtrToLastTuple:poPtrToLastTuple+InPagePointerSize], uint16Zero)

	page := Page{
		table: tableDef,
		page:  rawPage,
	}
	page.updateHash()

	return page
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

func (p Page) TupleCount() uint16 {
	lastPointer := p.lastTuplePtr()
	if lastPointer == 0 {
		return 0
	}

	firstPointer := poPtrToFirstTuple

	return (lastPointer-InPagePointer(firstPointer))/2 + 1
}

func (p Page) FreeSpace() uint16 {
	lastTuplePtr := p.lastTuplePtr()
	lastTuplePtrPtr := binary.LittleEndian.Uint16(
		p.page[poPrtToLastTuplePtr : poPrtToLastTuplePtr+InPagePointerSize])

	return lastTuplePtr - (lastTuplePtrPtr + InPagePointerSize - 1)
}

func (p Page) InsertTuple(tuple []byte) error {
	tupleLen := uint16(len(tuple))

	if tupleLen+InPagePointerSize > p.FreeSpace() {
		return errors.New(fmt.Sprintf("cannot fit tuple of size: %d to page with free space: %d",
			tupleLen, p.FreeSpace()))
	}

	// copy tuple
	lastTuplePtr := p.lastTuplePtr()
	newTuplePtr := lastTuplePtr - tupleLen
	copy(p.page[newTuplePtr:lastTuplePtr], tuple)

	// create new tuple pointer
	tupleLenBytes := make([]byte, 0, 2)
	binary.LittleEndian.PutUint16(tupleLenBytes, tupleLen)

	lastTuplePtrPtr := p.ptrToLastTuplePtr()
	newTuplePtrPtr := lastTuplePtrPtr + InPagePointerSize
	copy(p.page[newTuplePtrPtr:newTuplePtrPtr+InPagePointerSize], tupleLenBytes)

	// reassign last pointer pointer
	copy(p.page[poPrtToLastTuplePtr:poPrtToLastTuplePtr+InPagePointerSize],
		tupleLenBytes)

	return nil
}

func (p Page) UpdateTuple(tIndex TupleIndex, newTuple []byte) {
	if tIndex <= p.TupleCount() {
		panic(fmt.Sprintf("page does not contains tuple with index %d", tIndex))
	}

	tuplePtr := p.tuplePtr(tIndex)
	prevTuplePtr := p.tuplePtr(tIndex - 1)
	tuple := p.page[tuplePtr:prevTuplePtr]
	if len(tuple) != len(newTuple) {
		panic("When updating tuple it's len must me identical")
	}

	copy(tuple, newTuple)
}

func (p Page) lastTuplePtr() InPagePointer {
	pointerToLastPointer := bparse.Parse.UInt2(
		p.page[poPrtToLastTuplePtr : poPrtToLastTuplePtr+InPagePointerSize])

	if pointerToLastPointer == 0 {
		return 0
	}

	return bparse.Parse.UInt2(p.page[pointerToLastPointer:InPagePointerSize])
}

func (p Page) tupleEnd(index TupleIndex) InPagePointer {
	tuplePtr := p.tuplePtr(index)
	if tuplePtr == p.lastTuplePtr() {
		return Size
	}

	prevTuplePtr := p.tuplePtr(index - 1)
	return prevTuplePtr
}

func (p Page) tuplePtr(index TupleIndex) InPagePointer {
	ptrStart := poPtrToFirstTuple + InPagePointerSize*index
	return bparse.Parse.UInt2(p.page[ptrStart : ptrStart+InPagePointerSize])
}

func (p Page) ptrToLastTuplePtr() InPagePointer {
	return binary.LittleEndian.Uint16(
		p.page[poPtrToLastTuple : poPtrToLastTuple+InPagePointerSize])
}

func (p Page) updateHash() {
	hash := md5.Sum(p.page[poPageHash+pageHashLen:])
	copy(p.page[poPageHash:poPageHash+pageHashLen], hash[0:pageHashLen])
}

type InPagePointer = uint16

const InPagePointerSize = 2

// po - page offset to
const (
	poPageHash          = 0                        // offset to page hash
	poPrtToLastTuplePtr = poPageHash + pageHashLen // offset to pointer pointing to last tuple pointer
	poPtrToLastTuple    = poPrtToLastTuplePtr + 2  // offset to pointer pointing to last tuple start
	poPtrToFirstTuple   = poPtrToLastTuple + 2     // offset to first of many tuple pointers
)

const (
	pageHashLen = 8
)
