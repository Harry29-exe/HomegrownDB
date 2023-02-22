package internal

import (
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

func NewTupleBuilder(table reldef.TableRDefinition) OptimisticTupleBuilder {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))
	return OptimisticTupleBuilder{builder}
}

type OptimisticTupleBuilder struct {
	builder page.TupleBuilder
}

func (o OptimisticTupleBuilder) WriteValue(value rawtype.Value) {
	err := o.builder.WriteValue(value)
	if err != nil {
		log.Panicf("unexpected err: %s", err)
	}
}

func (o OptimisticTupleBuilder) Tuple(tx tx.Tx) page.Tuple {
	return o.builder.VolatileTuple(tx)
}
