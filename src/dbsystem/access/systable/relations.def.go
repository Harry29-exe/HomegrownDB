package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

func RelationsTableDef() table.RDefinition {
	columns := []column.WDef{
		relations.oid(),
		relations.relKind(),
		relations.fsmOID(),
		relations.vmOID(),
	}

	return newTableDef(
		RelationsName,
		HGRelationsOID,
		HGRelationsFsmOID,
		HGRelationsVmOID,
		columns,
	)
}

const (
	RelationsOrderOID column.Order = iota
	RelationsOrderRelKind
	RelationsOrderFsmOID
	RelationsOrderVmOID
)

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() column.WDef {
	return column.NewDefinition(
		"id",
		HGRelationsColOID,
		RelationsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relKind() column.WDef {
	return column.NewDefinition(
		"rel_kind",
		HGRelationsColRelKind,
		RelationsOrderRelKind,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) fsmOID() column.WDef {
	return column.NewDefinition(
		"fsm_oid",
		HGRelationsColFsmOID,
		RelationsOrderFsmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) vmOID() column.WDef {
	return column.NewDefinition(
		"vm_oid",
		HGRelationsColVmOID,
		RelationsOrderVmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

//todo add locks columns
