package buffer

import (
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/storage/pageio"
)

// todo change methods to operate on ArrayIndexes
type StdBuffer interface {
	ReadRPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error)
	ReadWPage(ownerID dbobj.OID, pageId page.Id, strategy rbm) (stdPage, error)

	ReleaseWPage(tag pageio.PageTag)
	ReleaseRPage(tag pageio.PageTag)
}

// rbm  read buffer mode
type rbm = uint8

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
	Tag            pageio.PageTag
	IsNew          bool
	IsReadFromDisc bool
}
