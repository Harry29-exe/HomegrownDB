package systable

import (
	"HomegrownDB/dbsystem/hgtype/inputtype"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var relationsDef = RelationsTableDef()

type RelationsOps struct {
}

func (o RelationsOps) TableAsRelationsRow(
	table table.Definition,
	tx tx.Tx,
	commands uint16,
) page.Tuple {
	builder := newTupleBuilder(relationsDef)

	builder.WriteValue(inputtype.ConvInt8Value(int64(table.OID())))
	builder.WriteValue(inputtype.ConvInt8Value(int64(relation.TypeTable)))
	builder.WriteValue(inputtype.ConvInt8Value(int64(table.FsmOID())))
	builder.WriteValue(inputtype.ConvInt8Value(int64(table.VmOID())))

	return builder.Tuple(tx, commands)
}
