package internal

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/bparse"
	"log"
)

func NewTableDef(name string, oid reldef.OID, fsmOID reldef.OID, vmOID reldef.OID, columns []column.WDef) tabdef.RDefinition {
	tableDef := tabdef.NewDefinition(name)
	tableDef.InitRel(oid, fsmOID, vmOID)

	for _, colDef := range columns {
		err := tableDef.AddColumn(colDef)
		if err != nil {
			log.Panicf("unexpected error while creating Relations tabdef: %s", err.Error())
		}
	}

	return tableDef
}

func NewColumnDef(name string, oid reldef.OID, order column.Order, ctype hgtype.ColType) column.WDef {
	return column.NewDefinition(name, oid, order, ctype)
}

// -------------------------
//      TupleBuilder
// -------------------------

func NewTupleBuilder(table tabdef.RDefinition) OptimisticTupleBuilder {
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

// -------------------------
//      DataConv
// -------------------------

func GetString(order column.Order, tuple page.RTuple) string {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return string(rawtype.TypeOp.GetData(value, args))
}

func GetInt8(order column.Order, tuple page.RTuple) int64 {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Int8(rawtype.TypeOp.GetData(value, args))
}

func GetBool(order column.Order, tuple page.RTuple) bool {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Bool(rawtype.TypeOp.GetData(value, args))
}
