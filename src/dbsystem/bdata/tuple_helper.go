package bdata

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
)

var TupleHelper = tupleHelper{}

type tupleHelper struct{}

func (th tupleHelper) GetValue(tuple Tuple, tableDef table.Definition, id column.Order) []byte {
	if tuple.IsNull(id) {
		return nil
	}
	if id == 0 {
		return tableDef.ColumnType(0).Value(tuple.Data())
	}
	remainingData := tuple.Data()
	columns := tableDef.Columns()

	for i := column.Order(0); i < id-1; i++ {
		if !tuple.IsNull(i) {
			remainingData = columns[i].CType().Skip(remainingData)
		}
	}

	return columns[id].CType().Value(remainingData)
}

func (th tupleHelper) GetValueByName(tuple Tuple, tableDef table.Definition, colName string) []byte {
	order, ok := tableDef.ColumnOrder(colName)
	if !ok {
		panic(fmt.Sprintf("No column with name %s on table %s", colName, tableDef.Name()))
	}

	return th.GetValue(tuple, tableDef, order)
}
