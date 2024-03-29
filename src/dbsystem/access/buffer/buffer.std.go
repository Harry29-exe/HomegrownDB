package buffer

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/storage/page"
)

// todo change methods to operate on ArrayIndexes
type StdBuffer interface {
	ReadRPage(ownerID hglib.OID, pageId page.Id, strategy rbm) (stdPage, error)
	ReadWPage(ownerID hglib.OID, pageId page.Id, strategy rbm) (stdPage, error)

	ReleaseWPage(tag page.PageTag)
	ReleaseRPage(tag page.PageTag)
	FlushAll() error
}

// rbm  read buffer mode
type rbm = uint8

type slotIndex = uint

const (
	// RbmReadOrCreate if page exist in buffer or in disc read it,
	// otherwise create page filled with zero in buffer
	RbmReadOrCreate rbm = iota
	// RbmRead if page exist return it otherwise return error
	RbmRead
	// RbmNoIO if page exist in buffer read it,
	// otherwise create new filled with zeros
	RbmNoIO
)

type stdPage struct {
	Bytes          []byte
	Tag            page.PageTag
	IsNew          bool
	IsReadFromDisc bool
}
