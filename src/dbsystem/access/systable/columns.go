package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
)

func ColumnsTableDef() table.RDefinition {
	tableCols := []column.WDef{
		columns.oid(),
		columns.relationOID(),
		columns.name(),
	}
	tableCols = append(tableCols, columns.argsColumns()...)

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

func (columnsBuilder) argsColumns() []column.WDef {
	return []column.WDef{
		column.NewDefinition(
			"arg_length",
			false,
			coltype.NewInt8(hgtype.Args{}),
		),
		//column.NewDefinition(
		//	"arg_nullable",
		//	false,
		//	hgtype.),
		//todo add more args when hgtype.bool will be implemented
	}
}
