package tpage

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/storage/page"
	"crypto/md5"
	"encoding/binary"
)

type InPagePointer = uint16

const InPagePointerSize = 2

var emptyPageFreeSpace = page.Size - (poFirstTuplePtr + InPagePointerSize)

func (p TablePage) getTupleEnd(index TupleIndex) InPagePointer {
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

func (p TablePage) getTupleStart(index TupleIndex) InPagePointer {
	ptrStart := poFirstTuplePtr + InPagePointerSize*index
	return bparse.Parse.UInt2(p.bytes[ptrStart : ptrStart+InPagePointerSize])
}

func (p TablePage) setTupleStart(tupleIndex TupleIndex, tupleStart InPagePointer) {
	ptrStart := poFirstTuplePtr + InPagePointerSize*tupleIndex
	binary.BigEndian.PutUint16(p.bytes[ptrStart:], tupleStart)
}

func (p TablePage) getPtrPosition(index TupleIndex) InPagePointer {
	return poFirstTuplePtr + index*InPagePointerSize
}

func (p TablePage) getLastPtrPosition() InPagePointer {
	return bparse.Parse.UInt2(
		p.bytes[poPrtToLastTuplePtr:])
}

func (p TablePage) setLastPointerPosition(ptr InPagePointer) {
	binary.BigEndian.PutUint16(p.bytes[poPrtToLastTuplePtr:], ptr)
}

func (p TablePage) getLastTupleStart() InPagePointer {
	return bparse.Parse.UInt2(p.bytes[poPtrToLastTupleStart:])
}

func (p TablePage) setLastTupleStart(ptr InPagePointer) {
	binary.BigEndian.PutUint16(p.bytes[poPtrToLastTupleStart:], ptr)
}

func (p TablePage) updateHash() {
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
