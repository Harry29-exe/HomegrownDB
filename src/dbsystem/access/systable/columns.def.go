package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
)

var columnsDef = createColumnsTableDef()

func ColumnsTableDef() reldef.TableRDefinition {
	return columnsDef
}

func createColumnsTableDef() reldef.TableRDefinition {
	colBuilder := columnsBuilder{}

	tableCols := []reldef.ColumnDefinition{
		colBuilder.oid(),
		colBuilder.relationOID(),
		colBuilder.colName(),
		colBuilder.colOrder(),
		colBuilder.typeTag(),
		colBuilder.argsLength(),
		colBuilder.argsNullable(),
		colBuilder.argsVarLen(),
		colBuilder.argsUTF8(),
	}

	return internal.NewTableDef(
		ColumnsName,
		HGColumnsOID,
		HGColumnsFsmOID,
		HGColumnsVmOID,
		tableCols,
	)
}

const (
	ColumnsOrderOID reldef.Order = iota
	ColumnsOrderRelationOID
	ColumnsOrderColName
	ColumnsOrderColOrder
	ColumnsOrderTypeTag
	ColumnsOrderArgsLength
	ColumnsOrderArgsNullable
	ColumnsOrderArgsVarLen
	ColumnsOrderArgsUTF8
)

type columnsBuilder struct{}

func (c *columnsBuilder) oid() (col reldef.ColumnDefinition) {
	return reldef.NewColumnDefinition(
		"id",
		HGColumnsColOID,
		ColumnsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) relationOID() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"relation_oid",
		HGColumnsColRelationOID,
		ColumnsOrderRelationOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) colName() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"col_name",
		HGColumnsColColName,
		ColumnsOrderColName,
		hgtype.NewStr(rawtype.Args{
			Length:   255,
			Nullable: false,
			VarLen:   true,
			UTF8:     false,
		}))
}

func (c *columnsBuilder) colOrder() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"col_order",
		HGColumnsColColOrder,
		ColumnsOrderColOrder,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) typeTag() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"type_tag",
		HGColumnsColTypeTag,
		ColumnsOrderTypeTag,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsLength() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"arg_length",
		HGColumnsColArgsLength,
		ColumnsOrderArgsLength,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsNullable() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"arg_nullable",
		HGColumnsColArgsNullable,
		ColumnsOrderArgsNullable,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsVarLen() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"arg_VarLen",
		HGColumnsColArgsVarLen,
		ColumnsOrderArgsVarLen,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsUTF8() reldef.ColumnDefinition {
	return reldef.NewColumnDefinition(
		"arg_UTF8",
		HGColumnsColArgsUTF8,
		ColumnsOrderArgsUTF8,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}
