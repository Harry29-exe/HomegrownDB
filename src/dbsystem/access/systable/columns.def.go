package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
)

var columnsDef = createColumnsTableDef()

func ColumnsTableDef() tabdef.RDefinition {
	return columnsDef
}

func createColumnsTableDef() tabdef.RDefinition {
	colBuilder := columnsBuilder{}

	tableCols := []tabdef.ColumnDefinition{
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
	ColumnsOrderOID tabdef.Order = iota
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

func (c *columnsBuilder) oid() (col tabdef.ColumnDefinition) {
	return tabdef.NewColumnDefinition(
		"id",
		HGColumnsColOID,
		ColumnsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) relationOID() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"relation_oid",
		HGColumnsColRelationOID,
		ColumnsOrderRelationOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) colName() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
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

func (c *columnsBuilder) colOrder() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"col_order",
		HGColumnsColColOrder,
		ColumnsOrderColOrder,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) typeTag() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"type_tag",
		HGColumnsColTypeTag,
		ColumnsOrderTypeTag,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsLength() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"arg_length",
		HGColumnsColArgsLength,
		ColumnsOrderArgsLength,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsNullable() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"arg_nullable",
		HGColumnsColArgsNullable,
		ColumnsOrderArgsNullable,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsVarLen() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"arg_VarLen",
		HGColumnsColArgsVarLen,
		ColumnsOrderArgsVarLen,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsUTF8() tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(
		"arg_UTF8",
		HGColumnsColArgsUTF8,
		ColumnsOrderArgsUTF8,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}
