package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
)

var relationsDef = createHGRelations()

func RelationsTableDef() tabdef.RDefinition {
	return relationsDef
}

func createHGRelations() tabdef.RDefinition {
	columns := []column.ColumnDefinition{
		relations.oid(),
		relations.relKind(),
		relations.relName(),
		relations.fsmOID(),
		relations.vmOID(),
	}

	return internal.NewTableDef(
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
	RelationsOrderRelName
	RelationsOrderFsmOID
	RelationsOrderVmOID
)

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() column.ColumnDefinition {
	return column.NewColumnDefinition(
		"id",
		HGRelationsColOID,
		RelationsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relKind() column.ColumnDefinition {
	return column.NewColumnDefinition(
		"rel_kind",
		HGRelationsColRelKind,
		RelationsOrderRelKind,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relName() column.ColumnDefinition {
	return column.NewColumnDefinition(
		"rel_name",
		HGRelationsColRelName,
		RelationsOrderRelName,
		hgtype.NewStr(rawtype.Args{Length: 255, Nullable: false, VarLen: true, UTF8: false}),
	)
}

func (relationsBuilder) fsmOID() column.ColumnDefinition {
	return column.NewColumnDefinition(
		"fsm_oid",
		HGRelationsColFsmOID,
		RelationsOrderFsmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) vmOID() column.ColumnDefinition {
	return column.NewColumnDefinition(
		"vm_oid",
		HGRelationsColVmOID,
		RelationsOrderVmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

//todo add locks columns
