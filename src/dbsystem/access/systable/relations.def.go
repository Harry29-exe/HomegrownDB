package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

type RelationsTable interface {
	table.RDefinition
	ColumnOID() column.Def
	ColumnRelKind() column.Def
	ColumnFsmOID() column.Def
	ColumnVmOID() column.Def
}

var _ RelationsTable = relationsTable{}

type relationsTable struct {
	table.RDefinition
}

func (r relationsTable) ColumnOID() column.Def {
	return r.Column(0)
}

func (r relationsTable) ColumnRelKind() column.Def {
	return r.Column(1)
}

func (r relationsTable) ColumnFsmOID() column.Def {
	return r.Column(2)
}

func (r relationsTable) ColumnVmOID() column.Def {
	return r.Column(3)
}

// -------------------------
//      Def
// -------------------------

func RelationsTableDef() RelationsTable {
	columns := []column.WDef{
		relations.oid(),
		relations.relKind(),
		relations.fsmOID(),
		relations.vmOID(),
	}

	return relationsTable{
		newTableDef(
			RelationsName,
			RelationsOID,
			RelationsFsmOID,
			RelationsVmOID,
			columns,
		),
	}
}

var relations = relationsBuilder{}

type relationsBuilder struct{}

func (relationsBuilder) oid() column.WDef {
	return column.NewDefinition(
		"id",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) relKind() column.WDef {
	return column.NewDefinition(
		"rel_kind",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) fsmOID() column.WDef {
	return column.NewDefinition(
		"fsm_oid",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (relationsBuilder) vmOID() column.WDef {
	return column.NewDefinition(
		"vm_oid",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

//todo add locks columns
