package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var RelationsOps = relationsOps{}

type relationsOps struct{}

func (relationsOps) TableAsRelationsRow(
	table tabdef.TableRDefinition,
	tx tx.Tx,
) (page.Tuple, error) {
	builder := internal.NewTupleBuilder(relationsDef)

	builder.WriteValue(intype.ConvInt8Value(int64(table.OID())))
	builder.WriteValue(intype.ConvInt8Value(int64(table.Kind())))
	name, err := intype.ConvStrValue(table.Name())
	if err != nil {
		return builder.Tuple(tx), err
	}
	builder.WriteValue(name)
	builder.WriteValue(intype.ConvInt8Value(int64(table.FsmOID())))
	builder.WriteValue(intype.ConvInt8Value(int64(table.VmOID())))

	return builder.Tuple(tx), nil
}

func (relationsOps) RowAsData(tuple page.Tuple) reldef.Relation {
	oid := internal.GetInt8(RelationsOrderOID, tuple)
	kind := reldef.Kind(internal.GetInt8(RelationsOrderRelKind, tuple))
	name := internal.GetString(RelationsOrderRelName, tuple)
	fsmOID := internal.GetInt8(RelationsOrderFsmOID, tuple)
	vmOID := internal.GetInt8(RelationsOrderVmOID, tuple)

	switch kind {
	case reldef.TypeTable:
		rel := tabdef.NewTableDefinition(name)
		rel.InitRel(reldef.OID(oid), reldef.OID(fsmOID), reldef.OID(vmOID))
		return rel
	default:
		//todo implement me
		panic("Not implemented")
	}
}
