package systable

import (
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
)

var relationsDef = RelationsTableDef()

type RelationsOps struct {
}

func (o RelationsOps) TableAsRelationsTuple(table table.Definition) page.Tuple {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))
	//builder.WriteValue(hgtype.Value{
	//	TypeTag:   inputtype.ConvInt8(0)),
	//	NormValue: nil,
	//}
	//todo implement me
	panic("Not implemented")
}
