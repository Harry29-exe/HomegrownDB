package systable

import (
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/storage/page"
	"log"
)

var columnsDef = ColumnsTableDef()

func ColumnAsColumnsRow(tableId OID, col column.Def) page.WTuple {
	builder := newTupleBuilder(columnsDef)

	builder.WriteValue(intype.ConvInt8Value(int64(col.Id())))
	builder.WriteValue(intype.ConvInt8Value(int64(tableId)))

	name, err := intype.ConvStrValue(col.Name())
	if err != nil {
		log.Panicf("enexpected err: %s", err.Error())
	}
	builder.WriteValue(name)
	//args
	//todo implement me
	panic("Not implemented")
}
