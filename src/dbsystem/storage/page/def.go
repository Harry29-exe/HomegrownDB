package page

import (
	"HomegrownDB/dbsystem/config"
	"HomegrownDB/dbsystem/relation/dbobj"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page/internal"
	"math"
)

type Id = uint32

const IdSize = 4

const (
	MaxId     = math.MaxUint32 - 1
	InvalidId = math.MaxUint32
)

const Size uint16 = config.PageSize

type PageTag = internal.PageTag

func NewTablePageTag(pageIndex Id, tableDef table.RDefinition) PageTag {
	return PageTag{
		PageId:  pageIndex,
		OwnerID: tableDef.OID(),
	}
}

func NewPageTag(pageIndex Id, objID dbobj.OID) PageTag {
	return PageTag{
		PageId:  pageIndex,
		OwnerID: objID,
	}
}
