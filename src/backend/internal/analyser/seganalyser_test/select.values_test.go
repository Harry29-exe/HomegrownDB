package seganalyser

import (
	"HomegrownDB/backend/internal/analyser"
	node "HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	testinfr "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/hgtype/rawtype"
	"HomegrownDB/lib/datastructs/appsync"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

func TestValuesSelect(t *testing.T) {
	//given
	inputQuery := "VALUES (1, 'Bob'), (2, 'Alice')"
	store := testinfr.NewRelationAccessManager()

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
	valuesRTE.ColTypes = []hgtype.ColumnType{
		hgtype.NewInt8(rawtype.Args{Nullable: true}),
		hgtype.NewStr(rawtype.Args{VarLen: true, Length: len("Alice"), UTF8: false, Nullable: true}),
	}
	query.RTables = []node.RangeTableEntry{valuesRTE}
	query.FromExpr = node.NewFromExpr2(nil, valuesRTE.CreateRef())
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(valuesRTE.Id, 0, valuesRTE.ColTypes[0]), 0, ""),
		node.NewTargetEntry(node.NewVar(valuesRTE.Id, 1, valuesRTE.ColTypes[1]), 1, ""),
	}
	return query
}
