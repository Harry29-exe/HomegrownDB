package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var columnsDef = ColumnsTableDef()

var Columns = columnsOps{}

type columnsOps struct{}

func (columnsOps) DataToRow(
	tableId OID,
	col column.Def,
	tx tx.Tx,
	commands uint16,
) page.WTuple {
	builder := newTupleBuilder(columnsDef)

	builder.WriteValue(intype.ConvInt8Value(int64(col.Id())))
	builder.WriteValue(intype.ConvInt8Value(int64(tableId)))

	name, err := intype.ConvStrValue(col.Name())
	if err != nil {
		log.Panicf("enexpected err: %s", err.Error())
	}
	builder.WriteValue(name)

	args := col.CType().Args()
	builder.WriteValue(intype.ConvInt8Value(int64(args.Length)))
	builder.WriteValue(intype.ConvBoolValue(args.Nullable))
	builder.WriteValue(intype.ConvBoolValue(args.VarLen))
	builder.WriteValue(intype.ConvBoolValue(args.UTF8))

	return builder.Tuple(tx, commands)
}

func (o columnsOps) DataToRows(
	tableOID OID,
	columns []column.Def,
	tx tx.Tx,
	commands uint16,
) []page.WTuple {
	tuples := make([]page.WTuple, len(columns))
	for i, colDef := range columns {
		tuples[i] = o.DataToRow(tableOID, colDef, tx, commands)
	}
	return tuples
}

func (o columnsOps) RowToData(row page.RTuple) column.WDef {
	name := GetString(ColumnsOrderColName, row)
	colOID := GetInt8(ColumnsOrderOID, row)
	order := GetInt8(ColumnsOrderColOrder, row)
	typeTag := GetInt8(ColumnsOrderTypeTag, row)
	argsLength := GetInt8(ColumnsOrderArgsLength, row)
	argsVarLen := GetBool(ColumnsOrderArgsVarLen, row)
	argsUTF8 := GetBool(ColumnsOrderArgsUTF8, row)
	argsNulable := GetBool(ColumnsOrderArgsNullable, row)

	return newColumnDef(
		name,
		OID(colOID),
		column.Order(order),
		hgtype.NewColType(rawtype.Tag(typeTag), hgtype.Args{
			Length:   int(argsLength),
			Nullable: argsNulable,
			VarLen:   argsVarLen,
			UTF8:     argsUTF8,
		}),
	)
}
