package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype/intype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var RelationsOps = relationsOps{}

type relationsOps struct{}

func (ops relationsOps) TableAsRelationsRow(
	table reldef.TableRDefinition,
	tx tx.Tx,
) (page.Tuple, error) {
	return ops.createRow(
		int64(table.OID()),
		int64(table.Kind()),
		table.Name(),
		int64(table.FsmOID()),
		int64(table.VmOID()),
		tx,
	)
}

func (ops relationsOps) SequenceAsRelationsRow(
	sequence reldef.SequenceDef,
	tx tx.Tx,
) (page.Tuple, error) {
	return ops.createRow(
		int64(sequence.OID()),
		int64(sequence.Kind()),
		sequence.Name(),
		int64(sequence.FsmOID()),
		int64(sequence.VmOID()),
		tx,
	)
}

func (ops relationsOps) createRow(
	oid int64,
	kind int64,
	name string,
	fsmOID int64,
	vmOID int64,
	tx tx.Tx,
) (page.Tuple, error) {
	builder := internal.NewTupleBuilder(relationsDef)
	builder.WriteValue(intype.ConvInt8Value(oid))
	builder.WriteValue(intype.ConvInt8Value(kind))
	nameVal, err := intype.ConvStrValue(name)
	if err != nil {
		return page.Tuple{}, err
	}
	builder.WriteValue(nameVal)
	builder.WriteValue(intype.ConvInt8Value(fsmOID))
	builder.WriteValue(intype.ConvInt8Value(vmOID))

	return builder.Tuple(tx), nil
}

func (relationsOps) RowAsData(tuple page.Tuple) reldef.BaseRelation {
	oid := internal.GetInt8(RelationsOrderOID, tuple)
	kind := reldef.Kind(internal.GetInt8(RelationsOrderRelKind, tuple))
	name := internal.GetString(RelationsOrderRelName, tuple)
	fsmOID := internal.GetInt8(RelationsOrderFsmOID, tuple)
	vmOID := internal.GetInt8(RelationsOrderVmOID, tuple)

	return reldef.BaseRelation{
		RelName:          name,
		ID:               reldef.OID(oid),
		RelKind:          kind,
		FreeSpaceMapOID:  reldef.OID(fsmOID),
		VisibilityMapOID: reldef.OID(vmOID),
	}
}
