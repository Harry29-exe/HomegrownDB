package internal

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/dbobj"
	"math"
)

type Id = uint32

const IdSize = 4

const (
	MaxId     = math.MaxUint32 - 1
	InvalidId = math.MaxUint32
)

const Size uint16 = config.PageSize

type PageTag struct {
	PageId  Id
	OwnerID dbobj.OID
}
