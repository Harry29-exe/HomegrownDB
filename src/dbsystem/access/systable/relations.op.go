package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/inputtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
)

var relationsDef = RelationsTableDef()

type RelationsOps struct {
}

func (o RelationsOps) TableAsRelationsTuple(table table.Definition) page.Tuple {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))
	builder.WriteValue(hgtype.Value{
		TypeTag:   inputtype.ConvInt8(0)),
		NormValue: nil,
	}
}
