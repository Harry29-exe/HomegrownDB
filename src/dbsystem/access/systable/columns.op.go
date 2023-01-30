package systable

import (
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var columnsDef = ColumnsTableDef()

func ColumnAsColumnsRow(
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

	args := col.CType().Args
	builder.WriteValue(intype.ConvInt8Value(int64(args.Length)))
	builder.WriteValue(intype.ConvBoolValue(args.Nullable))
	builder.WriteValue(intype.ConvBoolValue(args.VarLen))
	builder.WriteValue(intype.ConvBoolValue(args.UTF8))

	return builder.Tuple(tx, commands)
}
