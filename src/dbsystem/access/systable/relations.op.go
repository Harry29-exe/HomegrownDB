package systable

import (
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
)

type RelationsOps struct {
	Def table.RDefinition
}

func (o RelationsOps) TableAsRelationsTuple(table table.Definition) page.Tuple {
	//todo implement me
	panic("Not implemented")
	//tuple := page.NewTempTuple()
	//o.Def.Column(0).CType().WriteValue()
}
