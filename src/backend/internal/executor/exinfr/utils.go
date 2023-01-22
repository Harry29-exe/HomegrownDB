package exinfr

import (
	node "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/storage/dpage"
	"math"
)

func PatternFromTargetList(targetList []node.TargetEntry) *dpage.TuplePattern {
	pattern := &dpage.TuplePattern{
		Columns:   make([]dpage.ColumnInfo, len(targetList)),
		BitmapLen: uint16(math.Ceil(float64(len(targetList)) / 8)),
	}

	for i := 0; i < len(targetList); i++ {
		pattern.Columns[i] = typeFromTargetEntry(targetList[i])
	}

	return pattern
}

func typeFromTargetEntry(entry node.TargetEntry) dpage.ColumnInfo {
	entryType := entry.TypeTag().Type()

	return dpage.ColumnInfo{
		CType: entryType,
		Name:  entry.ColName,
	}
}

func PattenFromRTE(rte node.RangeTableEntry) *dpage.TuplePattern {
	switch rte.Kind {
	case node.RteValues:
		colTypes := make([]dpage.ColumnInfo, len(rte.ColTypes))
		for i := 0; i < len(rte.ColTypes); i++ {
			colTypes[i] = dpage.ColumnInfo{
				CType: rte.ColTypes[i].Type,
				Name:  rte.ColAlias[i].AliasName,
			}
		}
		return dpage.NewPattern(colTypes)
	case node.RteRelation:
		return dpage.NewPatternFromTable(rte.Ref)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
