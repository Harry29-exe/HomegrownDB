package pager

const PageSize uint32 = 8192

type Page struct {
	pageHash       int8
	lastRowPointer InPagePointer
	lastRowStart   InPagePointer
	rowPointers    []InPagePointer
	rows           []byte
}

type InPagePointer = uint16
