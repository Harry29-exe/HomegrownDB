package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
	"log"
)

func RelationsTableDef() table.RDefinition {
	tab := table.NewDefinition(RelationsName)
	tab.InitRel(RelationsOID, RelationsFsmOID, RelationsVmOID)
	columns := []column.WDef{
		relations.oid(),
		relations.relKind(),
		relations.fsmOID(),
		relations.vmOID(),
	}
	for _, colDef := range columns {
		err := tab.AddColumn(colDef)
		if err != nil {
			log.Panicf("unexpected error while creating Relations table: %s", err.Error())
		}
	}

	return tab
}

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (b relationsBuilder) oid() column.WDef {
	return column.NewDefinition(
		"id",
		false,
		hgtype.NewInt8(hgtype.Args{}),
	)
}

func (b relationsBuilder) relKind() column.WDef {
	return column.NewDefinition(
		"rel_kind",
		false,
		hgtype.NewInt8(hgtype.Args{}),
	)
}

func (b relationsBuilder) fsmOID() column.WDef {
	return column.NewDefinition(
		"fsm_oid",
		false,
		hgtype.NewInt8(hgtype.Args{}),
	)
}

func (b relationsBuilder) vmOID() column.WDef {
	return column.NewDefinition(
		"vm_oid",
		false,
		hgtype.NewInt8(hgtype.Args{}),
	)
}
