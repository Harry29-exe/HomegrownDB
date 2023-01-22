package exinfr

import (
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/page"
	"math"
)

func PatternFromTargetList(targetList []node.TargetEntry) page.TuplePattern {
	pattern := page.TuplePattern{
		Columns:   make([]page.ColumnInfo, len(targetList)),
		BitmapLen: uint16(math.Ceil(float64(len(targetList)) / 8)),
	}

	for i := 0; i < len(targetList); i++ {
		pattern.Columns[i] = typeFromTargetEntry(targetList[i])
	}

	return pattern
}

func typeFromTargetEntry(entry node.TargetEntry) page.ColumnInfo {
	entryType := entry.TypeTag().Type()

	return page.ColumnInfo{
		CType: entryType,
		Name:  entry.ColName,
	}
}

func PattenFromRTE(rte node.RangeTableEntry) page.TuplePattern {
	switch rte.Kind {
	case node.RteValues:
		colTypes := make([]page.ColumnInfo, len(rte.ColTypes))
		for i := 0; i < len(rte.ColTypes); i++ {
			colTypes[i] = page.ColumnInfo{
				CType: rte.ColTypes[i].Type,
				Name:  rte.ColAlias[i].AliasName,
			}
		}
		return page.NewPattern(colTypes)
	case node.RteRelation:
		return page.PatternFromTable(rte.Ref)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
