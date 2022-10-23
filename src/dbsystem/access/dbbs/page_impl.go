package dbbs

// For better understanding of struct Page
// it's recommended to view page.doc.svg diagram
// which shows how page's binary representations looks like

import (
	"HomegrownDB/dbsystem/schema/table"
	"encoding/binary"
	"errors"
	"fmt"
)

//todo add handling for inserting into empty page

func CreateEmptyPage(tableDef table.Definition) Page {
	rawPage := make([]byte, PageSize)
	uint16Zero := make([]byte, 2)
	binary.BigEndian.PutUint16(uint16Zero, 0)

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

func (p Page) Tuple(tIndex TupleIndex) Tuple {
	if p.TupleCount() <= tIndex {
		panic(fmt.Sprintf("Page has %d tuples but was requestd tuple with id: %d",
			p.TupleCount(), tIndex))
	}
	tuplePtr := p.getTupleStart(tIndex)
	tupleEndPtr := p.getTupleEnd(tIndex)

	return Tuple{
		data:  p.page[tuplePtr:tupleEndPtr],
		table: p.table,
	}
}

func (p Page) TupleCount() uint16 {
	lastPtrPosition := p.getLastPtrPosition()
	if lastPtrPosition == 0 {
		return 0
	}

	firstPtrIndex := poFirstTuplePtr

	return (lastPtrPosition-InPagePointer(firstPtrIndex))/2 + 1
}

func (p Page) FreeSpace() uint16 {
	if lastTupleStart := p.getLastTupleStart(); lastTupleStart != 0 {
		return lastTupleStart - (p.getLastPtrPosition() + InPagePointerSize)
	}
	return emptyPageFreeSpace
}

func (p Page) Data() []byte {
	return p.page
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
	var tuplePtrPosition InPagePointer

	insertedAsLast := true
	// empty page
	if lastTupleStart == 0 {
		lastTupleStart = PageSize
		tuplePtrPosition = poFirstTuplePtr
	} else {
		tCount := p.TupleCount()
		for i := uint16(0); i < tCount; i++ {
			if tupleStart := p.getTupleStart(i); tupleStart == 0 {
				tuplePtrPosition = poFirstTuplePtr + tuplePtrPosition*InPagePointerSize
				insertedAsLast = false
				break
			}
		}

		tuplePtrPosition = p.getLastPtrPosition() + InPagePointerSize
	}

	tupleStartIndex := lastTupleStart - tupleLen
	copy(p.page[tupleStartIndex:lastTupleStart], tuple)
	binary.BigEndian.PutUint16(p.page[tuplePtrPosition:], tupleStartIndex)
	p.setLastTupleStart(tupleStartIndex)
	if insertedAsLast {
		p.setLastPointerPosition(tuplePtrPosition)
	}

	return nil
}

func (p Page) UpdateTuple(tIndex TupleIndex, newTuple []byte) {
	if tIndex <= p.TupleCount() {
		panic(fmt.Sprintf("page does not contains tuple with index %d", tIndex))
	}

	tuplePtr := p.getTupleStart(tIndex)
	prevTuplePtr := p.getTupleStart(tIndex - 1)
	tuple := p.page[tuplePtr:prevTuplePtr]
	if len(tuple) != len(newTuple) {
		panic("When updating tuple it's len must me identical")
	}

	copy(tuple, newTuple)
}

func (p Page) DeleteTuple(tIndex TupleIndex) {
	tCount := p.TupleCount()
	if tCount == 0 || tCount <= tIndex {
		panic(fmt.Sprintf(
			"Page has %d tuples but DeleteTuple was invoked with parameter tIndex=%d",
			tCount, tIndex))
	}
	if tCount == 1 {
		p.setLastPointerPosition(0)
		p.setLastTupleStart(0)

	} else if tIndex == tCount-1 {
		newLastPtrIndex := p.TupleCount() - 2
		newLastTupleStart := p.getTupleStart(newLastPtrIndex)
		p.setLastTupleStart(newLastTupleStart)
		p.setLastPointerPosition(p.getPtrPosition(newLastPtrIndex))

	} else {
		p.deleteTupleFromMiddle(tIndex)
	}
}

func (p Page) deleteTupleFromMiddle(tIndex TupleIndex) {
	deletedTupleEnd := p.getTupleEnd(tIndex)
	deletedTupleStart := p.getTupleStart(tIndex)
	deletedTupleLen := deletedTupleEnd - deletedTupleStart
	lastTupleStart := p.getLastTupleStart()

	copy(p.page[lastTupleStart+deletedTupleLen:deletedTupleEnd], p.page[lastTupleStart:deletedTupleStart])
	tCount := p.TupleCount()
	for i := TupleIndex(0); i < tCount; i++ {
		if tupleStart := p.getTupleStart(i); tupleStart != 0 && tupleStart < deletedTupleStart {
			p.setTupleStart(i, tupleStart+deletedTupleLen)
		}
	}
	p.setLastTupleStart(p.getLastTupleStart() + deletedTupleLen)
	p.setTupleStart(tIndex, 0)
}
