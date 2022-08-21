package bdata

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"fmt"
)

var TupleHelper = tupleHelper{}

type tupleHelper struct{}

func (th tupleHelper) GetValue(tuple Tuple, tableDef table.Definition, id column.OrderId) []byte {
	if id == 0 {
		return tableDef.ColumnParser(0).GetValue(tuple.Data())
	}
	remainingData := tuple.Data()
	parsers := tableDef.AllColumnParsers()

	for i := column.OrderId(0); i < id-1; i++ {
		if !tuple.IsNull(i) {
			remainingData = parsers[i].Skip(remainingData)
		}
	}

	return parsers[id].GetValue(remainingData)
}

func (th tupleHelper) GetValueByName(tuple Tuple, tableDef table.Definition, colName string) []byte {
	order, ok := tableDef.ColumnId(colName)
	if !ok {
		panic(fmt.Sprintf("No column with name %s on table %s", colName, tableDef.Name()))
	}

	return th.GetValue(tuple, tableDef, order)
}
