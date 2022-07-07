package bdata

// For better understanding of struct Page
// it's recommended to view page.doc.svg diagram
// which shows how page's binary representations looks like

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/schema/table"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
)

//todo add handling for inserting into empty page

func CreateEmptyPage(tableDef table.Definition) Page {
	rawPage := make([]byte, PageSize)
	uint16Zero := make([]byte, 2)
	binary.LittleEndian.PutUint16(uint16Zero, 0)

	copy(rawPage[poPrtToLastTuplePtr:poPrtToLastTuplePtr+InPagePointerSize], uint16Zero)
	copy(rawPage[poPtrToLastTupleStart:poPtrToLastTupleStart+InPagePointerSize], uint16Zero)

	page := Page{
		table: tableDef,
		page:  rawPage,
	}
	page.updateHash()

	return page
}

func NewPage(definition table.Definition, data []byte) Page {
	return Page{definition, data}
}

type Page struct {
	table table.Definition
	page  []byte
}

type PageTag struct {
	PageId  PageId
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
	lastPtrIndex := p.getLastPtrIndex()
	if lastPtrIndex == 0 {
		return 0
	}

	firstPtrIndex := poFirstTuplePtr

	return (lastPtrIndex-InPagePointer(firstPtrIndex))/2 + 1
}

func (p Page) FreeSpace() uint16 {
	lastTupleStartIndex := p.getLastTupleStart()
	lastTuplePointerStart := p.getLastPtrIndex()

	return lastTupleStartIndex - (lastTuplePointerStart + InPagePointerSize - 1)
}

/*
todo add inserting tuple to page that have dead index eg. page had tuples 1, 2 and 3
 then tuple 2 was deleted next insert should put new tuple between 1 and 3 */

func (p Page) InsertTuple(tuple []byte) error {
	tupleLen := uint16(len(tuple))

	if tupleLen+InPagePointerSize > p.FreeSpace() {
		return errors.New(fmt.Sprintf("cannot fit tuple of size: %d to page with free space: %d",
			tupleLen, p.FreeSpace()))
	}

	// copy tuple
	lastTupleStart := p.getLastTupleStart()
	var tuplePtrIndex InPagePointer

	// empty page
	if lastTupleStart == 0 {
		lastTupleStart = PageSize
		tuplePtrIndex = poFirstTuplePtr
	} else {
		tuplePtrIndex = p.getPtrToLastTuplePtr() + InPagePointerSize
	}

	tupleStartIndex := lastTupleStart - tupleLen
	copy(p.page[tupleStartIndex:lastTupleStart], tuple)
	binary.LittleEndian.PutUint16(p.page[tuplePtrIndex:], tupleStartIndex)
	p.setPtrToLastTupleStart(tupleStartIndex)
	p.setPtrToLastTuplePtr(tuplePtrIndex)

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

func (p Page) DeleteTuple(tIndex TupleIndex) {
	//TODO implement me
	panic("implement me")
}

func (p Page) getLastTupleStart() InPagePointer {
	pointerToLastPointer := bparse.Parse.UInt2(p.page[poPrtToLastTuplePtr:])

	if pointerToLastPointer == 0 {
		return 0
	}

	return bparse.Parse.UInt2(p.page[pointerToLastPointer:])
}

func (p Page) tupleEnd(index TupleIndex) InPagePointer {
	tuplePtr := p.tuplePtr(index)
	if tuplePtr == p.getLastTupleStart() {
		return PageSize
	}

	prevTuplePtr := p.tuplePtr(index - 1)
	return prevTuplePtr
}

func (p Page) tuplePtr(index TupleIndex) InPagePointer {
	ptrStart := poFirstTuplePtr + InPagePointerSize*index
	return bparse.Parse.UInt2(p.page[ptrStart : ptrStart+InPagePointerSize])
}

func (p Page) getLastPtrIndex() InPagePointer {
	return binary.LittleEndian.Uint16(p.page[poPrtToLastTuplePtr:])
}

func (p Page) getPtrToLastTuplePtr() InPagePointer {
	return binary.LittleEndian.Uint16(
		p.page[poPtrToLastTupleStart:])
}

func (p Page) setPtrToLastTuplePtr(ptr InPagePointer) {
	binary.LittleEndian.PutUint16(p.page[poPrtToLastTuplePtr:], ptr)
}

func (p Page) getPtrToStartOfLastTuple() InPagePointer {
	return binary.LittleEndian.Uint16(p.page[poPtrToLastTupleStart:])
}

func (p Page) setPtrToLastTupleStart(ptr InPagePointer) {
	binary.LittleEndian.PutUint16(p.page[poPtrToLastTupleStart:], ptr)
}

func (p Page) updateHash() {
	hash := md5.Sum(p.page[poPageHash+pageHashLen:])
	copy(p.page[poPageHash:poPageHash+pageHashLen], hash[0:pageHashLen])
}

type InPagePointer = uint16

const InPagePointerSize = 2

// po - page offset to
const (
	poPageHash            = 0                         // offset to page hash
	poPrtToLastTuplePtr   = poPageHash + pageHashLen  // offset to pointer pointing to last tuple pointer
	poPtrToLastTupleStart = poPrtToLastTuplePtr + 2   // offset to pointer pointing to last tuple start
	poFirstTuplePtr       = poPtrToLastTupleStart + 2 // offset to first of many tuple pointers
)

const (
	pageHashLen = 8
)
