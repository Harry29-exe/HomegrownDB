package systable

import (
	"HomegrownDB/dbsystem/hgtype/inputtype"
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/storage/page"
	"log"
)

var columnsDef = ColumnsTableDef()

func ColumnAsColumnsRow(tableId OID, col column.Def) page.WTuple {
	builder := newTupleBuilder(columnsDef)

	builder.WriteValue(inputtype.ConvInt8Value(int64(col.Id())))
	builder.WriteValue(inputtype.ConvInt8Value(int64(tableId)))

	name, err := inputtype.ConvStrValue(col.Name())
	if err != nil {
		log.Panicf("enexpected err: %s", err.Error())
	}
	builder.WriteValue(name)
	//args
	//todo implement me
	panic("Not implemented")
}
