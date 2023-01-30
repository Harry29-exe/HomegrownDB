package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

type ColumnsTable interface {
	ColumnOID() column.Def
	ColumnRelationOID() column.Def
	ColumnName() column.Def
	ColumnsArgsLength() column.Def
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
		coltype.NewInt8(hgtype.Args{}),
	)
}

func (columnsBuilder) relationOID() column.WDef {
	return column.NewDefinition(
		"relation_oid",
		false,
		coltype.NewInt8(hgtype.Args{}),
	)
}

func (columnsBuilder) name() column.WDef {
	return column.NewDefinition(
		"col_name",
		false,
		coltype.NewStr(hgtype.Args{
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
		coltype.NewInt8(hgtype.Args{}),
	)
}
