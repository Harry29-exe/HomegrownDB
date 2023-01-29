package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

// -------------------------
//      Def
// -------------------------

func RelationsTableDef() table.RDefinition {
	columns := []column.WDef{
		relations.oid(),
		relations.relKind(),
		relations.fsmOID(),
		relations.vmOID(),
	}

	return newTableDef(
		RelationsName,
		RelationsOID,
		RelationsFsmOID,
		RelationsVmOID,
		columns,
	)
}

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() column.WDef {
	return column.NewDefinition(
		"id",
		false,
		coltype.NewInt8(hgtype.Args{}),
	)
}

func (relationsBuilder) relKind() column.WDef {
	return column.NewDefinition(
		"rel_kind",
		false,
		coltype.NewInt8(hgtype.Args{}),
	)
}

func (relationsBuilder) fsmOID() column.WDef {
	return column.NewDefinition(
		"fsm_oid",
		false,
		coltype.NewInt8(hgtype.Args{}),
	)
}

func (relationsBuilder) vmOID() column.WDef {
	return column.NewDefinition(
		"vm_oid",
		false,
		coltype.NewInt8(hgtype.Args{}),
	)
}

//todo add locks columns
