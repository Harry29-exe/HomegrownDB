package exinfr

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/storage/dpage"
	"math"
)

func PatternFromTargetList(targetList []node.TargetEntry) *dpage.TuplePattern {
	pattern := &dpage.TuplePattern{
		Columns:   make([]hgtype.TypeData, len(targetList)),
		BitmapLen: uint16(math.Ceil(float64(len(targetList)) / 8)),
	}

	for i := 0; i < len(targetList); i++ {
		pattern.Columns[i] = typeFromTargetEntry(targetList[i])
	}

	return pattern
}

func typeFromTargetEntry(entry node.TargetEntry) hgtype.TypeData {
	switch entry.ExprToExec.Tag() {
	case node.TagConst:
		return hgtype.NewTypeDataWithDefaultArgs(entry.Type())
	case node.TagVar:
		return entry.ExprToExec.(node.Var).TypeData
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func PattenFromRTE(rte node.RangeTableEntry) *dpage.TuplePattern {
	switch rte.Kind {
	case node.RteValues:
		colTypes := make([]hgtype.TypeData, len(rte.ColTypes))
		copy(colTypes, rte.ColTypes)
		return newPattern(colTypes)
	case node.RteRelation:
		relationCols := rte.Ref.Columns()
		colTypes := make([]hgtype.TypeData, len(relationCols))
		for i := 0; i < len(colTypes); i++ {
			colTypes[i] = relationCols[i].CType()
		}
		return newPattern(colTypes)
	default:
		//todo implement me
		panic("Not implemented")
	}
}

func newPattern(types []hgtype.TypeData) *dpage.TuplePattern {
	return &dpage.TuplePattern{
		Columns:   types,
		BitmapLen: uint16(math.Ceil(float64(len(types)) / 8)),
	}
}
