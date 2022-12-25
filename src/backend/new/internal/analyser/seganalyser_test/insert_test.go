package seganalyser

import (
	"HomegrownDB/backend/new/internal/analyser"
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/backend/new/internal/parser"
	. "HomegrownDB/backend/new/internal/testinfr"
	"HomegrownDB/common/datastructs/appsync"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/tt_user"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

func TestInsertAnalyse_SimplePositive1(t *testing.T) {
	//given
	query := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	store, users := TestTableStore.StoreWithUsersTable(t)
	expectedNode := InsertTests.expectedSimplePositive1(users, t)

	//when
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	queryNode, err := analyser.Analyse(pTree, store)
	assert.ErrIsNil(err, t)

	//then
	NodeAssert.Eq(expectedNode, queryNode, t)
}

func (insertTest) expectedSimplePositive1(users table.Definition, t *testing.T) node.Query {
	rteIdCounter := appsync.NewSimpleCounter[node.RteID](0)

	query := node.NewQuery(node.CommandTypeInsert, nil)
	query.TargetList = []node.TargetEntry{
		node.NewTargetEntry(nil, tt_user.C0IdOrder, tt_user.C0Id),
		node.NewTargetEntry(nil, tt_user.C1AgeOrder, tt_user.C1Age),
		node.NewTargetEntry(nil, tt_user.C2NameOrder, tt_user.C2Name),
	}

	resultRel := node.NewRelationRTE(rteIdCounter.Next(), users)
	query.RTables = []node.RangeTableEntry{resultRel}
	query.ResultRel = resultRel.Id

	subQuery := node.NewQuery(node.CommandTypeSelect, nil)
	valuesRte := node.NewValuesRTE(rteIdCounter.Next(), [][]node.Expr{
		{node.NewConstInt8(1), NewConstStr("bob", t)},
	})
	subQuery.RTables = []node.RangeTableEntry{valuesRte}
	subQuery.TargetList = []node.TargetEntry{
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 0, hgtype.NewInt8(hgtype.Args{})), 0, ""),
		node.NewTargetEntry(node.NewConst(hgtype.TypeInt8, nil), 1, ""),
		node.NewTargetEntry(node.NewVar(valuesRte.Id, 1, hgtype.NewStr(hgtype.Args{})), 0, ""),
	}
	subQuery.FromExpr = node.NewFromExpr2(nil, valuesRte.CreateRef())

	subQueryRte := node.NewSubqueryRTE(rteIdCounter.Next(), subQuery)
	query.RTables = append(query.RTables, subQueryRte)
	query.FromExpr = node.NewFromExpr2(nil, subQueryRte.CreateRef())

	return query
}

var InsertTests = insertTest{}

type insertTest struct{}
