package systable

import (
	"HomegrownDB/dbsystem/hgtype/inputtype"
	"HomegrownDB/dbsystem/relation"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

var relationsDef = RelationsTableDef()

type RelationsOps struct {
}

func (o RelationsOps) TableAsRelationsRow(
	table table.Definition,
	tx tx.Tx,
	commands uint16,
) page.Tuple {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))

	err := builder.WriteValue(inputtype.ConvInt8Value(int64(table.OID())))
	if err != nil {
		log.Printf("unexpected err: %s\n", err.Error())
	}
	err = builder.WriteValue(inputtype.ConvInt8Value(int64(relation.TypeTable)))
	if err != nil {
		log.Printf("unexpected err: %s\n", err.Error())
	}
	err = builder.WriteValue(inputtype.ConvInt8Value(int64(table.FsmOID())))
	if err != nil {
		log.Printf("unexpected err: %s\n", err.Error())
	}
	err = builder.WriteValue(inputtype.ConvInt8Value(int64(table.VmOID())))
	if err != nil {
		log.Printf("unexpected err: %s\n", err.Error())
	}

	return builder.VolatileTuple(tx, commands)
}
