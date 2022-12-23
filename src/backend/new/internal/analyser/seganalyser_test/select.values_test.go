package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	"HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/dbsystem/hgtype"
	"testing"
)

func TestValuesSelect(t *testing.T) {
	//given
	inputQuery := "VALUES (1, 'Bob'), (2, 'Alice')"
	store := testinfr.TestTableStore.TableStore(t)

	//when
	parseTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)
	query, err := analyser.Analyse(parseTree, store)
	assert.ErrIsNil(err, t)

	//then
	expectedQuery := SelectValues.expectedValuesSelect(t)
	testinfr.NodeAssert.Eq(expectedQuery, query, t)
}

var SelectValues = selectValues{}

type selectValues struct{}

func (selectValues) expectedValuesSelect(t *testing.T) node.Query {
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	query := node.NewQuery(node.CommandTypeSelect, nil)
	valuesRTE := node.NewValuesRTE(rteIdCounter.Next(), [][]node.Expr{
		{node.NewConstInt8(1), testinfr.NewConstStr("Bob", t)},
		{node.NewConstInt8(2), testinfr.NewConstStr("Alice", t)},
	})
	valuesRTE.ColTypes = []hgtype.TypeData{
		hgtype.NewInt8(hgtype.Args{}),
		hgtype.NewStr(hgtype.Args{VarLen: true, Length: uint32(len("Alice")), UTF8: false}),
	}
	query.RTables = []node.RangeTableEntry{valuesRTE}
	query.FromExpr = node.NewFromExpr2(nil, valuesRTE.CreateRef())
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(valuesRTE.Id, 0, valuesRTE.ColTypes[0]), 0, ""),
		node.NewTargetEntry(node.NewVar(valuesRTE.Id, 1, valuesRTE.ColTypes[1]), 1, ""),
	}
	return query
}
