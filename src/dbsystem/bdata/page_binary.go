package bdata

import (
	"HomegrownDB/common/bparse"
	"crypto/md5"
	"encoding/binary"
)

type InPagePointer = uint16

const InPagePointerSize = 2

var emptyPageFreeSpace = PageSize - (poFirstTuplePtr + InPagePointerSize)

func (p Page) tupleEnd(index TupleIndex) InPagePointer {
	if index == 0 {
		return PageSize
	}

	return p.tupleStart(index - 1)
}

func (p Page) tupleStart(index TupleIndex) InPagePointer {
	ptrStart := poFirstTuplePtr + InPagePointerSize*index
	return bparse.Parse.UInt2(p.page[ptrStart : ptrStart+InPagePointerSize])
}

func (p Page) getLastPtrPosition() InPagePointer {
	return bparse.Parse.UInt2(
		p.page[poPrtToLastTuplePtr:])
}

func (p Page) setLastPointerPosition(ptr InPagePointer) {
	binary.LittleEndian.PutUint16(p.page[poPrtToLastTuplePtr:], ptr)
}

func (p Page) getLastTupleStart() InPagePointer {
	return bparse.Parse.UInt2(p.page[poPtrToLastTupleStart:])
	//pointerToLastPointer := bparse.Parse.UInt2(p.page[poPrtToLastTuplePtr:])
	//
	//if pointerToLastPointer == 0 {
	//	return 0
	//}
	//
	//return bparse.Parse.UInt2(p.page[pointerToLastPointer:])
}

func (p Page) setLastTupleStart(ptr InPagePointer) {
	binary.LittleEndian.PutUint16(p.page[poPtrToLastTupleStart:], ptr)
}

func (p Page) updateHash() {
	hash := md5.Sum(p.page[poPageHash+pageHashLen:])
	copy(p.page[poPageHash:poPageHash+pageHashLen], hash[0:pageHashLen])
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
