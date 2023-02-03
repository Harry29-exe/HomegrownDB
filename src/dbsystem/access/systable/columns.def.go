package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

func ColumnsTableDef() table.RDefinition {
	colBuilder := columnsBuilder{}

	tableCols := []column.WDef{
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

	return newTableDef(
		ColumnsName,
		HGColumnsOID,
		HGColumnsFsmOID,
		HGColumnsVmOID,
		tableCols,
	)
}

const (
	ColumnsOrderOID column.Order = iota
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

func (c *columnsBuilder) oid() (col column.WDef) {
	return column.NewDefinition(
		"id",
		HGColumnsColOID,
		ColumnsOrderOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) relationOID() column.WDef {
	return column.NewDefinition(
		"relation_oid",
		HGColumnsColOID,
		ColumnsOrderRelationOID,
		hgtype.NewInt8(rawtype.Args{}),
	)
}

func (c *columnsBuilder) colName() column.WDef {
	return column.NewDefinition(
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

func (c *columnsBuilder) colOrder() column.WDef {
	return column.NewDefinition(
		"col_order",
		HGColumnsColColOrder,
		ColumnsOrderColOrder,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) typeTag() column.WDef {
	return column.NewDefinition(
		"type_tag",
		HGColumnsColTypeTag,
		ColumnsOrderTypeTag,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsLength() column.WDef {
	return column.NewDefinition(
		"arg_length",
		HGColumnsColArgsLength,
		ColumnsOrderArgsLength,
		hgtype.NewInt8(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsNullable() column.WDef {
	return column.NewDefinition(
		"arg_nullable",
		HGColumnsColArgsNullable,
		ColumnsOrderArgsNullable,
		hgtype.NewBool(hgtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsVarLen() column.WDef {
	return column.NewDefinition(
		"arg_VarLen",
		HGColumnsColArgsVarLen,
		ColumnsOrderArgsVarLen,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}

func (c *columnsBuilder) argsUTF8() column.WDef {
	return column.NewDefinition(
		"arg_UTF8",
		HGColumnsColArgsUTF8,
		ColumnsOrderArgsUTF8,
		hgtype.NewBool(rawtype.Args{Nullable: false}),
	)
}
