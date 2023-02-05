package internal

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/access/relation/table"
	"HomegrownDB/dbsystem/access/relation/table/column"
	"HomegrownDB/dbsystem/access/systable"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
	"log"
)

func NewTableDef(name string, oid systable.OID, fsmOID systable.OID, vmOID systable.OID, columns []column.WDef) table.RDefinition {
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

func NewColumnDef(name string, oid systable.OID, order column.Order, ctype hgtype.ColType) column.WDef {
	return column.NewDefinition(name, oid, order, ctype)
}

// -------------------------
//      TupleBuilder
// -------------------------

func NewTupleBuilder(table table.RDefinition) OptimisticTupleBuilder {
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
