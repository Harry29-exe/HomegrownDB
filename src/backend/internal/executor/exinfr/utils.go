package exinfr

import (
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/storage/dpage"
	"math"
)

func PatternFromTargetList(targetList []node2.TargetEntry) *dpage.TuplePattern {
	pattern := &dpage.TuplePattern{
		Columns:   make([]dpage.ColumnInfo, len(targetList)),
		BitmapLen: uint16(math.Ceil(float64(len(targetList)) / 8)),
	}

	for i := 0; i < len(targetList); i++ {
		pattern.Columns[i] = typeFromTargetEntry(targetList[i])
	}

	return pattern
}

func typeFromTargetEntry(entry node2.TargetEntry) dpage.ColumnInfo {
	var entryType hgtype.TypeData
	switch entry.ExprToExec.Tag() {
	case node2.TagConst:
		entryType = hgtype.NewTypeDataWithDefaultArgs(entry.Type())
	case node2.TagVar:
		entryType = entry.ExprToExec.(node2.Var).TypeData
	default:
		//todo implement me
		panic("Not implemented")
	}

	return dpage.ColumnInfo{
		CType: entryType,
		Name:  entry.ColName,
	}
}

func PattenFromRTE(rte node2.RangeTableEntry) *dpage.TuplePattern {
	switch rte.Kind {
	case node2.RteValues:
		colTypes := make([]dpage.ColumnInfo, len(rte.ColTypes))
		for i := 0; i < len(rte.ColTypes); i++ {
			colTypes[i] = dpage.ColumnInfo{
				CType: rte.ColTypes[i],
				Name:  rte.ColAlias[i].AliasName,
			}
		}
		return dpage.NewPattern(colTypes)
	case node2.RteRelation:
		return dpage.NewPatternFromTable(rte.Ref)
	default:
		//todo implement me
		panic("Not implemented")
	}
}
