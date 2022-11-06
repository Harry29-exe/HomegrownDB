package buffer

import (
	"HomegrownDB/dbsystem/schema/relation"
	"HomegrownDB/dbsystem/storage/page"
)

// todo change methods to operate on ArrayIndexes
type internalBuffer interface {
	ReadRPage(relation relation.Relation, pageId page.Id, strategy rbm) (buffPage, error)
	ReadWPage(relation relation.Relation, pageId page.Id, strategy rbm) (buffPage, error)

	ReleaseWPage(tag page.Tag)
	ReleaseRPage(tag page.Tag)
}

const NewPage = page.InvalidId

// rbm  read buffer mode
type rbm = uint8

const (
	// rbmReadOrCreate if page exist in buffer or in disc read it,
	// otherwise create page filled with zero in buffer
	rbmReadOrCreate rbm = iota
	// rbmRead if page exist return it otherwise return error
	rbmRead
	// rbmNoIO if page exist in buffer read it,
	// otherwise create new filled with zeros
	rbmNoIO
)

type buffPage struct {
	bytes          []byte
	tag            page.Tag
	isNew          bool
	isReadFromDisc bool
}
