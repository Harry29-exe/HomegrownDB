package seganalyser

import (
	"HomegrownDB/backend/internal/analyser"
	node2 "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	testinfr2 "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/coltype"
	"testing"
)

func TestValuesSelect(t *testing.T) {
	//given
	inputQuery := "VALUES (1, 'Bob'), (2, 'Alice')"
	store := testinfr2.TestTableStore.TableStore(t)

	//when
	parseTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)
	query, err := analyser.Analyse(parseTree, store)
	assert.ErrIsNil(err, t)

	//then
	expectedQuery := SelectValues.expectedValuesSelect(t)
	testinfr2.NodeAssert.Eq(expectedQuery, query, t)
}

var SelectValues = selectValues{}

type selectValues struct{}

func (selectValues) expectedValuesSelect(t *testing.T) node2.Query {
	rteIdCounter := appsync.NewSimpleCounter[node2.RteID](0)

	query := node2.NewQuery(node2.CommandTypeSelect, nil)
	valuesRTE := node2.NewValuesRTE(rteIdCounter.Next(), [][]node2.Expr{
		{node2.NewConstInt8(1), testinfr2.NewConstStr("Bob", t)},
		{node2.NewConstInt8(2), testinfr2.NewConstStr("Alice", t)},
	})
	valuesRTE.ColTypes = []coltype.ColumnType{
		coltype.NewInt8(hgtype.Args{}),
		coltype.NewStr(hgtype.Args{VarLen: true, Length: uint32(len("Alice")), UTF8: false}),
	}
	query.RTables = []node2.RangeTableEntry{valuesRTE}
	query.FromExpr = node2.NewFromExpr2(nil, valuesRTE.CreateRef())
	query.TargetList = []node2.TargetEntry{
		node2.NewTargetEntry(node2.NewVar(valuesRTE.Id, 0, valuesRTE.ColTypes[0]), 0, ""),
		node2.NewTargetEntry(node2.NewVar(valuesRTE.Id, 1, valuesRTE.ColTypes[1]), 1, ""),
	}
	return query
}
