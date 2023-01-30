package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

func newTableDef(name string, oid OID, fsmOID OID, vmOID OID, columns []column.WDef) table.RDefinition {
	tableDef := table.NewDefinition(name)
	tableDef.InitRel(oid, fsmOID, vmOID)

	for _, colDef := range columns {
		err := tableDef.AddColumn(colDef)
		if err != nil {
			log.Panicf("unexpected error while creating Relations table: %s", err.Error())
		}
	}

	return tableDef
}

func newTupleBuilder(table table.RDefinition) optimisticTupleBuilder {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))
	return optimisticTupleBuilder{builder}
}

type optimisticTupleBuilder struct {
	builder page.TupleBuilder
}

func (o optimisticTupleBuilder) WriteValue(value hgtype.Value) {
	err := o.builder.WriteValue(value)
	if err != nil {
		log.Panicf("unexpected err: %s", err)
	}
}

func (o optimisticTupleBuilder) Tuple(tx tx.Tx, commands uint16) page.Tuple {
	return o.builder.VolatileTuple(tx, commands)
}
