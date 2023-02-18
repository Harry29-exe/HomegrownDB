package internal

import (
	"HomegrownDB/dbsystem/hglib"
	"math"
)

type Id = uint32

const IdSize = 4

const (
	MaxId     = math.MaxUint32 - 1
	InvalidId = math.MaxUint32
)

const Size uint16 = 8192

type PageTag struct {
	PageId  Id
	OwnerID hglib.OID
}
