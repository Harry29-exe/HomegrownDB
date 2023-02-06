package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var relationsDef = RelationsTableDef()

var RelationsOps = relationsOps{}

type relationsOps struct{}

func (relationsOps) TableAsRelationsRow(
	table tabdef.RDefinition,
	tx tx.Tx,
) page.Tuple {
	builder := internal.NewTupleBuilder(relationsDef)

	builder.WriteValue(intype.ConvInt8Value(int64(table.OID())))
	builder.WriteValue(intype.ConvInt8Value(int64(table.Kind())))
	builder.WriteValue(intype.ConvInt8Value(int64(table.FsmOID())))
	builder.WriteValue(intype.ConvInt8Value(int64(table.VmOID())))

	return builder.Tuple(tx)
}
