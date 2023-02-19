package internal

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/bparse"
	"log"
)

func NewTableDef(name string, oid reldef.OID, fsmOID reldef.OID, vmOID reldef.OID, columns []tabdef.ColumnDefinition) tabdef.RDefinition {
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

func NewColumnDef(name string, oid reldef.OID, order tabdef.Order, ctype hgtype.ColType) tabdef.ColumnDefinition {
	return tabdef.NewColumnDefinition(name, oid, order, ctype)
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

func GetString(order tabdef.Order, tuple page.RTuple) string {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return string(rawtype.TypeOp.GetData(value, args))
}

func GetInt8(order tabdef.Order, tuple page.RTuple) int64 {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Int8(rawtype.TypeOp.GetData(value, args))
}

func GetBool(order tabdef.Order, tuple page.RTuple) bool {
	value := tuple.ColValue(order)
	args := tuple.Pattern().Columns[order].Type.Args()
	return bparse.Parse.Bool(rawtype.TypeOp.GetData(value, args))
}
