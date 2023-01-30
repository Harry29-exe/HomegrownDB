package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

type ColumnsTable interface {
	ColumnOID() column.Def
	ColumnRelationOID() column.Def
	ColumnName() column.Def
	ColumnArgsLength() column.Def
	ColumnArgsNullable() column.Def
	ColumnArgsVarLen() column.Def
	ColumnArgsUTF8() column.Def
}

var _ ColumnsTable = columnsTable{}

type columnsTable struct {
	table.RDefinition
}

func (c columnsTable) ColumnOID() column.Def {
	return c.RDefinition.Column(0)
}

func (c columnsTable) ColumnRelationOID() column.Def {
	return c.RDefinition.Column(1)
}

func (c columnsTable) ColumnName() column.Def {
	return c.RDefinition.Column(2)
}

func (c columnsTable) ColumnArgsLength() column.Def {
	return c.RDefinition.Column(3)
}

func (c columnsTable) ColumnArgsNullable() column.Def {
	return c.RDefinition.Column(4)
}

func (c columnsTable) ColumnArgsVarLen() column.Def {
	return c.RDefinition.Column(5)
}

func (c columnsTable) ColumnArgsUTF8() column.Def {
	return c.RDefinition.Column(6)
}

func ColumnsTableDef() table.RDefinition {
	tableCols := []column.WDef{
		columns.oid(),
		columns.relationOID(),
		columns.name(),
		columns.argsLength(),
	}

	return newTableDef(
		ColumnsName,
		ColumnsOID,
		ColumnsFsmOID,
		ColumnsVmOID,
		tableCols,
	)
}

var columns = columnsBuilder{}

type columnsBuilder struct{}

func (columnsBuilder) oid() column.WDef {
	return column.NewDefinition(
		"id",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (columnsBuilder) relationOID() column.WDef {
	return column.NewDefinition(
		"relation_oid",
		false,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (columnsBuilder) name() column.WDef {
	return column.NewDefinition(
		"col_name",
		false,
		hgtype.NewStr(rawtype.Args{
			Length:   255,
			Nullable: false,
			VarLen:   true,
			UTF8:     false,
		}))
}

func (columnsBuilder) argsLength() column.WDef {
	return column.NewDefinition(
		"arg_length",
		false,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (columnsBuilder) argsNullable() column.WDef {
	return column.NewDefinition(
		"arg_nullable",
		false,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

func (columnsBuilder) argsVarLen() column.WDef {
	return column.NewDefinition(
		"arg_VarLen",
		false,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}

func (columnsBuilder) argsUTF8() column.WDef {
	return column.NewDefinition(
		"arg_UTF8",
		false,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}
