package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var ColumnsOps = columnsOps{}

type columnsOps struct{}

func (columnsOps) DataToRow(tableId OID, col column.Def, tx tx.Tx) page.WTuple {
	builder := internal.NewTupleBuilder(columnsDef)

	builder.WriteValue(intype.ConvInt8Value(int64(col.Id())))
	builder.WriteValue(intype.ConvInt8Value(int64(tableId)))

	name, err := intype.ConvStrValue(col.Name())
	if err != nil {
		log.Panicf("enexpected err: %s", err.Error())
	}
	builder.WriteValue(name)
	builder.WriteValue(intype.ConvInt8Value(int64(col.Order())))
	builder.WriteValue(intype.ConvInt8Value(int64(col.CType().Tag())))

	args := col.CType().Args()
	builder.WriteValue(intype.ConvInt8Value(int64(args.Length)))
	builder.WriteValue(intype.ConvBoolValue(args.Nullable))
	builder.WriteValue(intype.ConvBoolValue(args.VarLen))
	builder.WriteValue(intype.ConvBoolValue(args.UTF8))

	return builder.Tuple(tx)
}

func (o columnsOps) DataToRows(tableOID OID, columns []column.Def, tx tx.Tx) []page.WTuple {
	tuples := make([]page.WTuple, len(columns))
	for i, colDef := range columns {
		tuples[i] = o.DataToRow(tableOID, colDef, tx)
	}
	return tuples
}

func (o columnsOps) RowToData(row page.RTuple) column.WDef {
	name := internal.GetString(ColumnsOrderColName, row)
	colOID := internal.GetInt8(ColumnsOrderOID, row)
	order := internal.GetInt8(ColumnsOrderColOrder, row)
	typeTag := internal.GetInt8(ColumnsOrderTypeTag, row)
	argsLength := internal.GetInt8(ColumnsOrderArgsLength, row)
	argsVarLen := internal.GetBool(ColumnsOrderArgsVarLen, row)
	argsUTF8 := internal.GetBool(ColumnsOrderArgsUTF8, row)
	argsNullable := internal.GetBool(ColumnsOrderArgsNullable, row)

	return internal.NewColumnDef(
		name,
		OID(colOID),
		column.Order(order),
		hgtype.NewColType(rawtype.Tag(typeTag), hgtype.Args{
			Length:   int(argsLength),
			Nullable: argsNullable,
			VarLen:   argsVarLen,
			UTF8:     argsUTF8,
		}),
	)
}

func (o columnsOps) TableOID(tuple page.RTuple) reldef.OID {
	return reldef.OID(internal.GetInt8(ColumnsOrderRelationOID, tuple))
}
