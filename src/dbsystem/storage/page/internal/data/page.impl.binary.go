package data

import (
	page "HomegrownDB/dbsystem/storage/page/internal"
	"HomegrownDB/lib/bparse"
	"crypto/md5"
	"encoding/binary"
)

type InPagePointer = uint16

const InPagePointerSize = 2

var emptyPageFreeSpace = page.Size - (poFirstTuplePtr + InPagePointerSize)

func (p Page) getTupleEnd(index TupleIndex) InPagePointer {
	var tupleEnd InPagePointer
	for index > 0 {
		tupleEnd = p.getTupleStart(index - 1)
		if tupleEnd != 0 {
			return tupleEnd
		}
		index--
	}

	return page.Size
}

func (p Page) getTupleStart(index TupleIndex) InPagePointer {
	ptrStart := poFirstTuplePtr + InPagePointerSize*index
	return bparse.Parse.UInt2(p.bytes[ptrStart : ptrStart+InPagePointerSize])
}

func (p Page) setTupleStart(tupleIndex TupleIndex, tupleStart InPagePointer) {
	ptrStart := poFirstTuplePtr + InPagePointerSize*tupleIndex
	binary.BigEndian.PutUint16(p.bytes[ptrStart:], tupleStart)
}

func (p Page) getPtrPosition(index TupleIndex) InPagePointer {
	return poFirstTuplePtr + index*InPagePointerSize
}

func (p Page) getLastPtrPosition() InPagePointer {
	return bparse.Parse.UInt2(
		p.bytes[poPrtToLastTuplePtr:])
}

func (p Page) setLastPointerPosition(ptr InPagePointer) {
	binary.BigEndian.PutUint16(p.bytes[poPrtToLastTuplePtr:], ptr)
}

func (p Page) getLastTupleStart() InPagePointer {
	return bparse.Parse.UInt2(p.bytes[poPtrToLastTupleStart:])
}

func (p Page) setLastTupleStart(ptr InPagePointer) {
	binary.BigEndian.PutUint16(p.bytes[poPtrToLastTupleStart:], ptr)
}

func (p Page) updateHash() {
	hash := md5.Sum(p.bytes[poPageHash+pageHashLen:])
	copy(p.bytes[poPageHash:poPageHash+pageHashLen], hash[0:pageHashLen])
}

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
