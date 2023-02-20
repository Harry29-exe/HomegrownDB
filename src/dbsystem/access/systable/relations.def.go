package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
)

var relationsDef = createHGRelations()

func RelationsTableDef() reldef.TableRDefinition {
	return relationsDef
}

func createHGRelations() reldef.TableRDefinition {
	columns := []reldef.ColumnDefinition{
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
	RelationsOrderOID reldef.Order = iota
	RelationsOrderRelKind
	RelationsOrderRelName
	RelationsOrderFsmOID
	RelationsOrderVmOID
)

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"id",
		HGRelationsColOID,
		RelationsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relKind() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"rel_kind",
		HGRelationsColRelKind,
		RelationsOrderRelKind,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relName() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"rel_name",
		HGRelationsColRelName,
		RelationsOrderRelName,
		hgtype.NewStr(rawtype.Args{Length: 255, Nullable: false, VarLen: true, UTF8: false}),
	)
}

func (relationsBuilder) fsmOID() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"fsm_oid",
		HGRelationsColFsmOID,
		RelationsOrderFsmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) vmOID() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"vm_oid",
		HGRelationsColVmOID,
		RelationsOrderVmOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

//todo add locks columns
