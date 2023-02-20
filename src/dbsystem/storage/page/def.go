package page

import (
	"HomegrownDB/dbsystem/hglib"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page/internal"
	"math"
)

type Id = uint32

const IdSize = 4

const (
	MaxId     = math.MaxUint32 - 1
	InvalidId = math.MaxUint32
)

const Size uint16 = internal.Size

type PageTag = internal.PageTag

func NewTablePageTag(pageIndex Id, tableDef reldef.TableRDefinition) PageTag {
	return PageTag{
		PageId:  pageIndex,
		OwnerID: tableDef.OID(),
	}
}

func NewPageTag(pageIndex Id, objID hglib.OID) PageTag {
	return PageTag{
		PageId:  pageIndex,
		OwnerID: objID,
	}
}
