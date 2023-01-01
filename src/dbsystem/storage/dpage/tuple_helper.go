package dpage

import (
	"HomegrownDB/dbsystem/relation/table"
	"HomegrownDB/dbsystem/relation/table/column"
	"fmt"
)

var TupleHelper = tupleHelper{}

type tupleHelper struct{}

func (th tupleHelper) GetValue(tuple Tuple, tableDef table.RDefinition, id column.Order) []byte {
	if tuple.IsNull(id) {
		return nil
	}
	if id == 0 {
		return tableDef.ColumnType(0).Value(tuple.Bytes())
	}
	remainingData := tuple.Bytes()
	columns := tableDef.Columns()

	for i := column.Order(0); i < id-1; i++ {
		if !tuple.IsNull(i) {
			remainingData = columns[i].CType().Skip(remainingData)
		}
	}

	return columns[id].CType().Value(remainingData)
}

func (th tupleHelper) GetValueByName(tuple Tuple, tableDef table.RDefinition, colName string) []byte {
	order, ok := tableDef.ColumnOrder(colName)
	if !ok {
		panic(fmt.Sprintf("No column with name %s on table %s", colName, tableDef.Name()))
	}

	return th.GetValue(tuple, tableDef, order)
}
