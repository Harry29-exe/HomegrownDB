package schema

type ID = uint32

type Relation interface {
	RelationID() ID
	PageInfo() PageInfo
}

type PageInfo struct {
	HeaderSize uint
	DataSize   uint
}
