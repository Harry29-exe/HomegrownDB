package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

var relationsDef = createHGRelations()

func RelationsTableDef() tabdef.TableRDefinition {
	return relationsDef
}

func createHGRelations() tabdef.TableRDefinition {
	columns := []tabdef.ColumnDefinition{
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
	RelationsOrderOID tabdef.Order = iota
	RelationsOrderRelKind
	RelationsOrderRelName
	RelationsOrderFsmOID
	RelationsOrderVmOID
)

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"id",
		HGRelationsColOID,
		RelationsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relKind() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"rel_kind",
		HGRelationsColRelKind,
		RelationsOrderRelKind,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relName() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"rel_name",
		HGRelationsColRelName,
		RelationsOrderRelName,
		hgtype.NewStr(rawtype.Args{Length: 255, Nullable: false, VarLen: true, UTF8: false}),
	)
}

func (relationsBuilder) fsmOID() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"fsm_oid",
		HGRelationsColFsmOID,
		RelationsOrderFsmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) vmOID() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"vm_oid",
		HGRelationsColVmOID,
		RelationsOrderVmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

//todo add locks columns
