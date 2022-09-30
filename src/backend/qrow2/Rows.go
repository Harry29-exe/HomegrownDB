package qrow2

import "HomegrownDB/dbsystem/ctype"

type Rows struct {
	R       []Row
	Pattern []ctype.CType
}

type Row struct {
	data    []byte
	Pattern []ctype.CType
}
