package systable

import (
	"HomegrownDB/common/bparse"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
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

func newColumnDef(name string, oid OID, order column.Order, ctype hgtype.ColType) column.WDef {
	return column.NewDefinition(name, oid, order, ctype)
}

// -------------------------
//      TupleBuilder
// -------------------------

func newTupleBuilder(table table.RDefinition) optimisticTupleBuilder {
	builder := page.NewTupleBuilder()
	builder.Init(page.PatternFromTable(table))
	return optimisticTupleBuilder{builder}
}

type optimisticTupleBuilder struct {
	builder page.TupleBuilder
}

func (o optimisticTupleBuilder) WriteValue(value rawtype.Value) {
	err := o.builder.WriteValue(value)
	if err != nil {
		log.Panicf("unexpected err: %s", err)
	}
}

func (o optimisticTupleBuilder) Tuple(tx tx.Tx, commands uint16) page.Tuple {
	return o.builder.VolatileTuple(tx, commands)
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
